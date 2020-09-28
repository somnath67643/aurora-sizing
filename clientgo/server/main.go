package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/somnath67643/aurora-sizing/clientgo/baseproto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterObjServiceService(grpcServer, &pb.ObjServiceService{
		Connect:            connect,
		LookupByNameStream: lookupByNameStream,
	})
	grpcServer.Serve(lis)
}

func connect(ctx context.Context, in *pb.CmdConnect) (*pb.CmdConnectResponse, error) {
	log.Println(in.AppName)
	resp := &pb.CmdConnectResponse{
		MsgSize:     111,
		VerAndRev:   222,
		FeatureFlag: 333,
	}
	return resp, nil
}

func lookupByNameStream(in *pb.CmdLookupByName, stream pb.ObjService_LookupByNameStreamServer) error {
	for i := 0; i < int(in.Count); i++ {
		resp := &pb.CmdLookupByNameResponseStream{
			SecurityName: fmt.Sprintf("%s_%d", in.SecurityNamePrefix, i),
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
	return nil
}