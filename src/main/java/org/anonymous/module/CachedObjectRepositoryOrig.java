package org.anonymous.module;

import io.prometheus.client.Counter;
import io.prometheus.client.Gauge;
import org.anonymous.domain.ObjectDTO;
import org.anonymous.grpc.CmdGetByNameExtResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisPool;
import redis.clients.jedis.JedisPoolConfig;

import java.util.List;
import java.util.Optional;

public class CachedObjectRepositoryOrig implements AutoCloseable {

    private static final Logger LOGGER = LoggerFactory.getLogger(CachedObjectRepositoryOrig.class);

    private static final Gauge getObjFromCacheGaugeTimer = Gauge.build().name("get_object_cache").help("Get Object on Middleware from cache").labelNames("redis").register();
    private static final Gauge getObjFromDBGaugeTimer = Gauge.build().name("get_object_db").help("Get Object on Middleware from db").labelNames("db").register();
    private static final Gauge setObjToCacheGaugeTimer = Gauge.build().name("set_object_cache").help("Set Object to cache").labelNames("cache").register();
    private static final Counter cacheOps = Counter.build().name("get_object_cache_count").help("Count of GetObject from Cache").labelNames("redis").register();
    private static final Counter dbOps = Counter.build().name("get_object_db_count").help("Count of GetObject from DB").labelNames("db").register();
    private static final Counter jedisConns = Counter.build().name("redis_connection_count").help("Count of Redis Connections").labelNames("redis").register();
    public static final int MAX_TOTAL = 5000;
    public static final String OBJ_MAP = "objMap";

    private final ObjectRepository delegate;
    private final JedisPool replicaPool;
    private final JedisPool primaryPool;
    private final ThreadLocal<Jedis> jedisROConnection;
    private final ThreadLocal<Jedis> jedisRWConnection;

    /*private static RedissonClient redisson;
    private static RLocalCachedMap<String, ObjectDTO> objMap;*/

    public CachedObjectRepositoryOrig(ObjectRepository objectRepository) {
        this.delegate = objectRepository;
        // max pool size and end point externalize
        // src/redis-cli -c -h aurora-sizing.uga7qd.ng.0001.use1.cache.amazonaws.com -p 6379

        JedisPoolConfig poolConfig = new JedisPoolConfig();
        poolConfig.setMaxTotal(MAX_TOTAL);
        poolConfig.setMaxWaitMillis(Integer.MAX_VALUE);
        this.primaryPool = new JedisPool(poolConfig, System.getProperty("redis.pri"));
        this.replicaPool = new JedisPool(poolConfig, System.getProperty("redis.rep"));
        jedisROConnection = ThreadLocal.withInitial(() -> {
            jedisConns.labels("conn count").inc();
            return replicaPool.getResource();
        });
        jedisRWConnection = ThreadLocal.withInitial(() -> {
            jedisConns.labels("conn count").inc();
            return primaryPool.getResource();
        });

        /*Config config = new Config();
        config.setTransportMode(TransportMode.EPOLL).setNettyThreads(500);
        config.useReplicatedServers()
                .addNodeAddress("redis://aurora-sizing-001.uga7qd.0001.use1.cache.amazonaws.com:6379")
                .addNodeAddress("redis://rep-1.uga7qd.0001.use1.cache.amazonaws.com:6379")
                .addNodeAddress("redis://rep-2.uga7qd.0001.use1.cache.amazonaws.com:6379")
                .addNodeAddress("redis://rep-3.uga7qd.0001.use1.cache.amazonaws.com:6379")
                .addNodeAddress("redis://rep-4.uga7qd.0001.use1.cache.amazonaws.com:6379")
                .setSlaveConnectionPoolSize(1000)
                .setMasterConnectionPoolSize(1000);

        redisson = Redisson.create(config);
        objMap = redisson.getLocalCachedMap(OBJ_MAP, LocalCachedMapOptions.defaults());*/

    }

    @Override
    public void close() throws Exception {
        primaryPool.close();
        replicaPool.close();
    }

    public List<String> lookup(String prefix, int typeid, int limit) {
        return delegate.lookup(prefix, typeid, limit);
    }

    public Optional<CmdGetByNameExtResponse.MsgOnSuccess> getFullObject(String key) {
        Optional<ObjectDTO> objectDTO = null;
        Gauge.Timer cacheTimer = getObjFromCacheGaugeTimer.labels("get_object_cache").startTimer();
        ObjectDTO fromCache = getFromRedis(key);
        if (null != fromCache) {
            cacheTimer.setDuration();
            cacheOps.labels("get_object_cache").inc();
            objectDTO = Optional.of(fromCache);
        } else {
            cacheTimer.close();
            Gauge.Timer dbTimer = getObjFromDBGaugeTimer.labels("get_object_db").startTimer();
            objectDTO = delegate.getFullObject(key);
            if (objectDTO.isPresent()) {
                dbTimer.setDuration();
                setToRedis(key, objectDTO.get());
                dbOps.labels("get_object_db").inc();
            } else {
                dbTimer.close();
            }
        }
        return Optional.ofNullable(objectDTO.get().toCmdGetByNameExtResponseMsgOnSuccess());
    }

    private void setToRedis(String key, ObjectDTO objectDTO) {
        Gauge.Timer setCacheTimer = setObjToCacheGaugeTimer.labels("set_object_cache").startTimer();
        try {
            jedisRWConnection.get().hset(OBJ_MAP.getBytes(), key.getBytes(), objectDTO.toBytes());
        } catch (Throwable th) {
            LOGGER.error("unexpected err in Redis set ", th);
        }
        setCacheTimer.setDuration();
    }

    private ObjectDTO getFromRedis(String key) {
        ObjectDTO objectDTO = null;
        try {
            byte[] bytes = jedisROConnection.get().hget(OBJ_MAP.getBytes(), key.getBytes());
            if ( null != bytes ) {
                objectDTO = ObjectDTO.fromBytes(bytes);
            }
        } catch (Throwable th) {
            LOGGER.error("unexpected err in Redis get ", th);
        }
        return objectDTO;
    }
}
