// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: cmd_get_by_name.proto

package org.anonymous.grpc;

public interface CmdGetByNameResponseOrBuilder extends
    // @@protoc_insertion_point(interface_extends:org.anonymous.grpc.CmdGetByNameResponse)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>int32 msg_size = 1;</code>
   */
  int getMsgSize();

  /**
   * <code>int32 status = 2;</code>
   */
  int getStatus();

  /**
   * <code>.org.anonymous.grpc.ErrorType error_type = 3;</code>
   */
  int getErrorTypeValue();
  /**
   * <code>.org.anonymous.grpc.ErrorType error_type = 3;</code>
   */
  org.anonymous.grpc.ErrorType getErrorType();

  /**
   * <code>bytes security = 4;</code>
   */
  com.google.protobuf.ByteString getSecurity();

  public org.anonymous.grpc.CmdGetByNameResponse.ResponseCase getResponseCase();
}
