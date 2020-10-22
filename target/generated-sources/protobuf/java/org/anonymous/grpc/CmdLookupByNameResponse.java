// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: cmd_lookup_by_name.proto

package org.anonymous.grpc;

/**
 * Protobuf type {@code org.anonymous.grpc.CmdLookupByNameResponse}
 */
public  final class CmdLookupByNameResponse extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:org.anonymous.grpc.CmdLookupByNameResponse)
    CmdLookupByNameResponseOrBuilder {
private static final long serialVersionUID = 0L;
  // Use CmdLookupByNameResponse.newBuilder() to construct.
  private CmdLookupByNameResponse(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private CmdLookupByNameResponse() {
    securityNames_ = com.google.protobuf.LazyStringArrayList.EMPTY;
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private CmdLookupByNameResponse(
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

            messageSize_ = input.readInt32();
            break;
          }
          case 16: {

            securityCount_ = input.readInt32();
            break;
          }
          case 26: {
            java.lang.String s = input.readStringRequireUtf8();
            if (!((mutable_bitField0_ & 0x00000004) != 0)) {
              securityNames_ = new com.google.protobuf.LazyStringArrayList();
              mutable_bitField0_ |= 0x00000004;
            }
            securityNames_.add(s);
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
      if (((mutable_bitField0_ & 0x00000004) != 0)) {
        securityNames_ = securityNames_.getUnmodifiableView();
      }
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.anonymous.grpc.CmdLookupByNameOuterClass.internal_static_org_anonymous_grpc_CmdLookupByNameResponse_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.anonymous.grpc.CmdLookupByNameOuterClass.internal_static_org_anonymous_grpc_CmdLookupByNameResponse_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.anonymous.grpc.CmdLookupByNameResponse.class, org.anonymous.grpc.CmdLookupByNameResponse.Builder.class);
  }

  private int bitField0_;
  public static final int MESSAGE_SIZE_FIELD_NUMBER = 1;
  private int messageSize_;
  /**
   * <code>int32 message_size = 1;</code>
   */
  public int getMessageSize() {
    return messageSize_;
  }

  public static final int SECURITY_COUNT_FIELD_NUMBER = 2;
  private int securityCount_;
  /**
   * <code>int32 security_count = 2;</code>
   */
  public int getSecurityCount() {
    return securityCount_;
  }

  public static final int SECURITY_NAMES_FIELD_NUMBER = 3;
  private com.google.protobuf.LazyStringList securityNames_;
  /**
   * <code>repeated string security_names = 3;</code>
   */
  public com.google.protobuf.ProtocolStringList
      getSecurityNamesList() {
    return securityNames_;
  }
  /**
   * <code>repeated string security_names = 3;</code>
   */
  public int getSecurityNamesCount() {
    return securityNames_.size();
  }
  /**
   * <code>repeated string security_names = 3;</code>
   */
  public java.lang.String getSecurityNames(int index) {
    return securityNames_.get(index);
  }
  /**
   * <code>repeated string security_names = 3;</code>
   */
  public com.google.protobuf.ByteString
      getSecurityNamesBytes(int index) {
    return securityNames_.getByteString(index);
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
    if (messageSize_ != 0) {
      output.writeInt32(1, messageSize_);
    }
    if (securityCount_ != 0) {
      output.writeInt32(2, securityCount_);
    }
    for (int i = 0; i < securityNames_.size(); i++) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 3, securityNames_.getRaw(i));
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (messageSize_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(1, messageSize_);
    }
    if (securityCount_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(2, securityCount_);
    }
    {
      int dataSize = 0;
      for (int i = 0; i < securityNames_.size(); i++) {
        dataSize += computeStringSizeNoTag(securityNames_.getRaw(i));
      }
      size += dataSize;
      size += 1 * getSecurityNamesList().size();
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
    if (!(obj instanceof org.anonymous.grpc.CmdLookupByNameResponse)) {
      return super.equals(obj);
    }
    org.anonymous.grpc.CmdLookupByNameResponse other = (org.anonymous.grpc.CmdLookupByNameResponse) obj;

    if (getMessageSize()
        != other.getMessageSize()) return false;
    if (getSecurityCount()
        != other.getSecurityCount()) return false;
    if (!getSecurityNamesList()
        .equals(other.getSecurityNamesList())) return false;
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
    hash = (37 * hash) + MESSAGE_SIZE_FIELD_NUMBER;
    hash = (53 * hash) + getMessageSize();
    hash = (37 * hash) + SECURITY_COUNT_FIELD_NUMBER;
    hash = (53 * hash) + getSecurityCount();
    if (getSecurityNamesCount() > 0) {
      hash = (37 * hash) + SECURITY_NAMES_FIELD_NUMBER;
      hash = (53 * hash) + getSecurityNamesList().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.anonymous.grpc.CmdLookupByNameResponse parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.anonymous.grpc.CmdLookupByNameResponse parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.anonymous.grpc.CmdLookupByNameResponse parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.anonymous.grpc.CmdLookupByNameResponse parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.anonymous.grpc.CmdLookupByNameResponse parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.anonymous.grpc.CmdLookupByNameResponse parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.anonymous.grpc.CmdLookupByNameResponse parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.anonymous.grpc.CmdLookupByNameResponse parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.anonymous.grpc.CmdLookupByNameResponse parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.anonymous.grpc.CmdLookupByNameResponse parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.anonymous.grpc.CmdLookupByNameResponse parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.anonymous.grpc.CmdLookupByNameResponse parseFrom(
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
  public static Builder newBuilder(org.anonymous.grpc.CmdLookupByNameResponse prototype) {
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
   * Protobuf type {@code org.anonymous.grpc.CmdLookupByNameResponse}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:org.anonymous.grpc.CmdLookupByNameResponse)
      org.anonymous.grpc.CmdLookupByNameResponseOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.anonymous.grpc.CmdLookupByNameOuterClass.internal_static_org_anonymous_grpc_CmdLookupByNameResponse_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.anonymous.grpc.CmdLookupByNameOuterClass.internal_static_org_anonymous_grpc_CmdLookupByNameResponse_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.anonymous.grpc.CmdLookupByNameResponse.class, org.anonymous.grpc.CmdLookupByNameResponse.Builder.class);
    }

    // Construct using org.anonymous.grpc.CmdLookupByNameResponse.newBuilder()
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
      messageSize_ = 0;

      securityCount_ = 0;

      securityNames_ = com.google.protobuf.LazyStringArrayList.EMPTY;
      bitField0_ = (bitField0_ & ~0x00000004);
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.anonymous.grpc.CmdLookupByNameOuterClass.internal_static_org_anonymous_grpc_CmdLookupByNameResponse_descriptor;
    }

    @java.lang.Override
    public org.anonymous.grpc.CmdLookupByNameResponse getDefaultInstanceForType() {
      return org.anonymous.grpc.CmdLookupByNameResponse.getDefaultInstance();
    }

    @java.lang.Override
    public org.anonymous.grpc.CmdLookupByNameResponse build() {
      org.anonymous.grpc.CmdLookupByNameResponse result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.anonymous.grpc.CmdLookupByNameResponse buildPartial() {
      org.anonymous.grpc.CmdLookupByNameResponse result = new org.anonymous.grpc.CmdLookupByNameResponse(this);
      int from_bitField0_ = bitField0_;
      int to_bitField0_ = 0;
      result.messageSize_ = messageSize_;
      result.securityCount_ = securityCount_;
      if (((bitField0_ & 0x00000004) != 0)) {
        securityNames_ = securityNames_.getUnmodifiableView();
        bitField0_ = (bitField0_ & ~0x00000004);
      }
      result.securityNames_ = securityNames_;
      result.bitField0_ = to_bitField0_;
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
      if (other instanceof org.anonymous.grpc.CmdLookupByNameResponse) {
        return mergeFrom((org.anonymous.grpc.CmdLookupByNameResponse)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.anonymous.grpc.CmdLookupByNameResponse other) {
      if (other == org.anonymous.grpc.CmdLookupByNameResponse.getDefaultInstance()) return this;
      if (other.getMessageSize() != 0) {
        setMessageSize(other.getMessageSize());
      }
      if (other.getSecurityCount() != 0) {
        setSecurityCount(other.getSecurityCount());
      }
      if (!other.securityNames_.isEmpty()) {
        if (securityNames_.isEmpty()) {
          securityNames_ = other.securityNames_;
          bitField0_ = (bitField0_ & ~0x00000004);
        } else {
          ensureSecurityNamesIsMutable();
          securityNames_.addAll(other.securityNames_);
        }
        onChanged();
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
      org.anonymous.grpc.CmdLookupByNameResponse parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.anonymous.grpc.CmdLookupByNameResponse) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }
    private int bitField0_;

    private int messageSize_ ;
    /**
     * <code>int32 message_size = 1;</code>
     */
    public int getMessageSize() {
      return messageSize_;
    }
    /**
     * <code>int32 message_size = 1;</code>
     */
    public Builder setMessageSize(int value) {
      
      messageSize_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 message_size = 1;</code>
     */
    public Builder clearMessageSize() {
      
      messageSize_ = 0;
      onChanged();
      return this;
    }

    private int securityCount_ ;
    /**
     * <code>int32 security_count = 2;</code>
     */
    public int getSecurityCount() {
      return securityCount_;
    }
    /**
     * <code>int32 security_count = 2;</code>
     */
    public Builder setSecurityCount(int value) {
      
      securityCount_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 security_count = 2;</code>
     */
    public Builder clearSecurityCount() {
      
      securityCount_ = 0;
      onChanged();
      return this;
    }

    private com.google.protobuf.LazyStringList securityNames_ = com.google.protobuf.LazyStringArrayList.EMPTY;
    private void ensureSecurityNamesIsMutable() {
      if (!((bitField0_ & 0x00000004) != 0)) {
        securityNames_ = new com.google.protobuf.LazyStringArrayList(securityNames_);
        bitField0_ |= 0x00000004;
       }
    }
    /**
     * <code>repeated string security_names = 3;</code>
     */
    public com.google.protobuf.ProtocolStringList
        getSecurityNamesList() {
      return securityNames_.getUnmodifiableView();
    }
    /**
     * <code>repeated string security_names = 3;</code>
     */
    public int getSecurityNamesCount() {
      return securityNames_.size();
    }
    /**
     * <code>repeated string security_names = 3;</code>
     */
    public java.lang.String getSecurityNames(int index) {
      return securityNames_.get(index);
    }
    /**
     * <code>repeated string security_names = 3;</code>
     */
    public com.google.protobuf.ByteString
        getSecurityNamesBytes(int index) {
      return securityNames_.getByteString(index);
    }
    /**
     * <code>repeated string security_names = 3;</code>
     */
    public Builder setSecurityNames(
        int index, java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  ensureSecurityNamesIsMutable();
      securityNames_.set(index, value);
      onChanged();
      return this;
    }
    /**
     * <code>repeated string security_names = 3;</code>
     */
    public Builder addSecurityNames(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  ensureSecurityNamesIsMutable();
      securityNames_.add(value);
      onChanged();
      return this;
    }
    /**
     * <code>repeated string security_names = 3;</code>
     */
    public Builder addAllSecurityNames(
        java.lang.Iterable<java.lang.String> values) {
      ensureSecurityNamesIsMutable();
      com.google.protobuf.AbstractMessageLite.Builder.addAll(
          values, securityNames_);
      onChanged();
      return this;
    }
    /**
     * <code>repeated string security_names = 3;</code>
     */
    public Builder clearSecurityNames() {
      securityNames_ = com.google.protobuf.LazyStringArrayList.EMPTY;
      bitField0_ = (bitField0_ & ~0x00000004);
      onChanged();
      return this;
    }
    /**
     * <code>repeated string security_names = 3;</code>
     */
    public Builder addSecurityNamesBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      ensureSecurityNamesIsMutable();
      securityNames_.add(value);
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


    // @@protoc_insertion_point(builder_scope:org.anonymous.grpc.CmdLookupByNameResponse)
  }

  // @@protoc_insertion_point(class_scope:org.anonymous.grpc.CmdLookupByNameResponse)
  private static final org.anonymous.grpc.CmdLookupByNameResponse DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.anonymous.grpc.CmdLookupByNameResponse();
  }

  public static org.anonymous.grpc.CmdLookupByNameResponse getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<CmdLookupByNameResponse>
      PARSER = new com.google.protobuf.AbstractParser<CmdLookupByNameResponse>() {
    @java.lang.Override
    public CmdLookupByNameResponse parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new CmdLookupByNameResponse(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<CmdLookupByNameResponse> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<CmdLookupByNameResponse> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.anonymous.grpc.CmdLookupByNameResponse getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

