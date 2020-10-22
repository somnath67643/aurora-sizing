// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: cmd_connect_ext.proto

package org.anonymous.grpc;

/**
 * Protobuf type {@code org.anonymous.grpc.CmdConnectExtResponse}
 */
public  final class CmdConnectExtResponse extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:org.anonymous.grpc.CmdConnectExtResponse)
    CmdConnectExtResponseOrBuilder {
private static final long serialVersionUID = 0L;
  // Use CmdConnectExtResponse.newBuilder() to construct.
  private CmdConnectExtResponse(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private CmdConnectExtResponse() {
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private CmdConnectExtResponse(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    int mutable_bitField0_ = 0;
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 8: {

            msgSize_ = input.readUInt32();
            break;
          }
          case 16: {

            verAndRev_ = input.readInt32();
            break;
          }
          case 24: {

            featureFlag_ = input.readUInt32();
            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.anonymous.grpc.CmdConnectExtOuterClass.internal_static_org_anonymous_grpc_CmdConnectExtResponse_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.anonymous.grpc.CmdConnectExtOuterClass.internal_static_org_anonymous_grpc_CmdConnectExtResponse_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.anonymous.grpc.CmdConnectExtResponse.class, org.anonymous.grpc.CmdConnectExtResponse.Builder.class);
  }

  public static final int MSG_SIZE_FIELD_NUMBER = 1;
  private int msgSize_;
  /**
   * <code>uint32 msg_size = 1;</code>
   */
  public int getMsgSize() {
    return msgSize_;
  }

  public static final int VER_AND_REV_FIELD_NUMBER = 2;
  private int verAndRev_;
  /**
   * <code>int32 ver_and_rev = 2;</code>
   */
  public int getVerAndRev() {
    return verAndRev_;
  }

  public static final int FEATURE_FLAG_FIELD_NUMBER = 3;
  private int featureFlag_;
  /**
   * <code>uint32 feature_flag = 3;</code>
   */
  public int getFeatureFlag() {
    return featureFlag_;
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (msgSize_ != 0) {
      output.writeUInt32(1, msgSize_);
    }
    if (verAndRev_ != 0) {
      output.writeInt32(2, verAndRev_);
    }
    if (featureFlag_ != 0) {
      output.writeUInt32(3, featureFlag_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (msgSize_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeUInt32Size(1, msgSize_);
    }
    if (verAndRev_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(2, verAndRev_);
    }
    if (featureFlag_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeUInt32Size(3, featureFlag_);
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof org.anonymous.grpc.CmdConnectExtResponse)) {
      return super.equals(obj);
    }
    org.anonymous.grpc.CmdConnectExtResponse other = (org.anonymous.grpc.CmdConnectExtResponse) obj;

    if (getMsgSize()
        != other.getMsgSize()) return false;
    if (getVerAndRev()
        != other.getVerAndRev()) return false;
    if (getFeatureFlag()
        != other.getFeatureFlag()) return false;
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + MSG_SIZE_FIELD_NUMBER;
    hash = (53 * hash) + getMsgSize();
    hash = (37 * hash) + VER_AND_REV_FIELD_NUMBER;
    hash = (53 * hash) + getVerAndRev();
    hash = (37 * hash) + FEATURE_FLAG_FIELD_NUMBER;
    hash = (53 * hash) + getFeatureFlag();
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.anonymous.grpc.CmdConnectExtResponse parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.anonymous.grpc.CmdConnectExtResponse parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.anonymous.grpc.CmdConnectExtResponse parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.anonymous.grpc.CmdConnectExtResponse parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.anonymous.grpc.CmdConnectExtResponse parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.anonymous.grpc.CmdConnectExtResponse parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.anonymous.grpc.CmdConnectExtResponse parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.anonymous.grpc.CmdConnectExtResponse parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.anonymous.grpc.CmdConnectExtResponse parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.anonymous.grpc.CmdConnectExtResponse parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.anonymous.grpc.CmdConnectExtResponse parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.anonymous.grpc.CmdConnectExtResponse parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(org.anonymous.grpc.CmdConnectExtResponse prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code org.anonymous.grpc.CmdConnectExtResponse}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:org.anonymous.grpc.CmdConnectExtResponse)
      org.anonymous.grpc.CmdConnectExtResponseOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.anonymous.grpc.CmdConnectExtOuterClass.internal_static_org_anonymous_grpc_CmdConnectExtResponse_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.anonymous.grpc.CmdConnectExtOuterClass.internal_static_org_anonymous_grpc_CmdConnectExtResponse_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.anonymous.grpc.CmdConnectExtResponse.class, org.anonymous.grpc.CmdConnectExtResponse.Builder.class);
    }

    // Construct using org.anonymous.grpc.CmdConnectExtResponse.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      msgSize_ = 0;

      verAndRev_ = 0;

      featureFlag_ = 0;

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.anonymous.grpc.CmdConnectExtOuterClass.internal_static_org_anonymous_grpc_CmdConnectExtResponse_descriptor;
    }

    @java.lang.Override
    public org.anonymous.grpc.CmdConnectExtResponse getDefaultInstanceForType() {
      return org.anonymous.grpc.CmdConnectExtResponse.getDefaultInstance();
    }

    @java.lang.Override
    public org.anonymous.grpc.CmdConnectExtResponse build() {
      org.anonymous.grpc.CmdConnectExtResponse result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.anonymous.grpc.CmdConnectExtResponse buildPartial() {
      org.anonymous.grpc.CmdConnectExtResponse result = new org.anonymous.grpc.CmdConnectExtResponse(this);
      result.msgSize_ = msgSize_;
      result.verAndRev_ = verAndRev_;
      result.featureFlag_ = featureFlag_;
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof org.anonymous.grpc.CmdConnectExtResponse) {
        return mergeFrom((org.anonymous.grpc.CmdConnectExtResponse)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.anonymous.grpc.CmdConnectExtResponse other) {
      if (other == org.anonymous.grpc.CmdConnectExtResponse.getDefaultInstance()) return this;
      if (other.getMsgSize() != 0) {
        setMsgSize(other.getMsgSize());
      }
      if (other.getVerAndRev() != 0) {
        setVerAndRev(other.getVerAndRev());
      }
      if (other.getFeatureFlag() != 0) {
        setFeatureFlag(other.getFeatureFlag());
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      org.anonymous.grpc.CmdConnectExtResponse parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.anonymous.grpc.CmdConnectExtResponse) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private int msgSize_ ;
    /**
     * <code>uint32 msg_size = 1;</code>
     */
    public int getMsgSize() {
      return msgSize_;
    }
    /**
     * <code>uint32 msg_size = 1;</code>
     */
    public Builder setMsgSize(int value) {
      
      msgSize_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>uint32 msg_size = 1;</code>
     */
    public Builder clearMsgSize() {
      
      msgSize_ = 0;
      onChanged();
      return this;
    }

    private int verAndRev_ ;
    /**
     * <code>int32 ver_and_rev = 2;</code>
     */
    public int getVerAndRev() {
      return verAndRev_;
    }
    /**
     * <code>int32 ver_and_rev = 2;</code>
     */
    public Builder setVerAndRev(int value) {
      
      verAndRev_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 ver_and_rev = 2;</code>
     */
    public Builder clearVerAndRev() {
      
      verAndRev_ = 0;
      onChanged();
      return this;
    }

    private int featureFlag_ ;
    /**
     * <code>uint32 feature_flag = 3;</code>
     */
    public int getFeatureFlag() {
      return featureFlag_;
    }
    /**
     * <code>uint32 feature_flag = 3;</code>
     */
    public Builder setFeatureFlag(int value) {
      
      featureFlag_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>uint32 feature_flag = 3;</code>
     */
    public Builder clearFeatureFlag() {
      
      featureFlag_ = 0;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:org.anonymous.grpc.CmdConnectExtResponse)
  }

  // @@protoc_insertion_point(class_scope:org.anonymous.grpc.CmdConnectExtResponse)
  private static final org.anonymous.grpc.CmdConnectExtResponse DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.anonymous.grpc.CmdConnectExtResponse();
  }

  public static org.anonymous.grpc.CmdConnectExtResponse getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<CmdConnectExtResponse>
      PARSER = new com.google.protobuf.AbstractParser<CmdConnectExtResponse>() {
    @java.lang.Override
    public CmdConnectExtResponse parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new CmdConnectExtResponse(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<CmdConnectExtResponse> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<CmdConnectExtResponse> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.anonymous.grpc.CmdConnectExtResponse getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

