// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: cmd_get_by_name_ext.proto

package org.anonymous.grpc;

public final class CmdGetByNameExtOuterClass {
  private CmdGetByNameExtOuterClass() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_org_anonymous_grpc_CmdGetByNameExt_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_org_anonymous_grpc_CmdGetByNameExt_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_org_anonymous_grpc_CmdGetByNameExtResponse_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_org_anonymous_grpc_CmdGetByNameExtResponse_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_org_anonymous_grpc_CmdGetByNameExtResponse_MsgOnSuccess_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_org_anonymous_grpc_CmdGetByNameExtResponse_MsgOnSuccess_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_org_anonymous_grpc_CmdGetByNameExtResponse_MsgOnFailure_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_org_anonymous_grpc_CmdGetByNameExtResponse_MsgOnFailure_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\031cmd_get_by_name_ext.proto\022\022org.anonymo" +
      "us.grpc\032\020error_type.proto\032\016metadata.prot" +
      "o\032\016cmd_type.proto\"\201\001\n\017CmdGetByNameExt\022\020\n" +
      "\010msg_size\030\001 \001(\005\022-\n\010msg_type\030\002 \001(\0162\033.org." +
      "anonymous.grpc.CmdType\022\026\n\016argument_flags" +
      "\030\003 \001(\005\022\025\n\rsecurity_name\030\004 \001(\t\"\270\003\n\027CmdGet" +
      "ByNameExtResponse\022R\n\016msg_on_success\030\002 \001(" +
      "\01328.org.anonymous.grpc.CmdGetByNameExtRe" +
      "sponse.MsgOnSuccessH\000\022R\n\016msg_on_failure\030" +
      "\003 \001(\01328.org.anonymous.grpc.CmdGetByNameE" +
      "xtResponse.MsgOnFailureH\000\032t\n\014MsgOnSucces" +
      "s\022\025\n\rhas_succeeded\030\001 \001(\010\022.\n\010metadata\030\002 \001" +
      "(\0132\034.org.anonymous.grpc.Metadata\022\020\n\010mem_" +
      "size\030\003 \001(\r\022\013\n\003mem\030\004 \001(\014\032l\n\014MsgOnFailure\022" +
      "\025\n\rhas_succeeded\030\001 \001(\010\0221\n\nerror_type\030\002 \001" +
      "(\0162\035.org.anonymous.grpc.ErrorType\022\022\n\nerr" +
      "or_text\030\003 \001(\tB\021\n\017RequestResponseB<P\001Z8gi" +
      "thub.com/somnath67643/aurora-sizing/clie" +
      "ntgo/baseprotob\006proto3"
    };
    com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner =
        new com.google.protobuf.Descriptors.FileDescriptor.    InternalDescriptorAssigner() {
          public com.google.protobuf.ExtensionRegistry assignDescriptors(
              com.google.protobuf.Descriptors.FileDescriptor root) {
            descriptor = root;
            return null;
          }
        };
    com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.anonymous.grpc.ErrorTypeOuterClass.getDescriptor(),
          org.anonymous.grpc.MetadataOuterClass.getDescriptor(),
          org.anonymous.grpc.CmdTypeOuterClass.getDescriptor(),
        }, assigner);
    internal_static_org_anonymous_grpc_CmdGetByNameExt_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_org_anonymous_grpc_CmdGetByNameExt_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_org_anonymous_grpc_CmdGetByNameExt_descriptor,
        new java.lang.String[] { "MsgSize", "MsgType", "ArgumentFlags", "SecurityName", });
    internal_static_org_anonymous_grpc_CmdGetByNameExtResponse_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_org_anonymous_grpc_CmdGetByNameExtResponse_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_org_anonymous_grpc_CmdGetByNameExtResponse_descriptor,
        new java.lang.String[] { "MsgOnSuccess", "MsgOnFailure", "RequestResponse", });
    internal_static_org_anonymous_grpc_CmdGetByNameExtResponse_MsgOnSuccess_descriptor =
      internal_static_org_anonymous_grpc_CmdGetByNameExtResponse_descriptor.getNestedTypes().get(0);
    internal_static_org_anonymous_grpc_CmdGetByNameExtResponse_MsgOnSuccess_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_org_anonymous_grpc_CmdGetByNameExtResponse_MsgOnSuccess_descriptor,
        new java.lang.String[] { "HasSucceeded", "Metadata", "MemSize", "Mem", });
    internal_static_org_anonymous_grpc_CmdGetByNameExtResponse_MsgOnFailure_descriptor =
      internal_static_org_anonymous_grpc_CmdGetByNameExtResponse_descriptor.getNestedTypes().get(1);
    internal_static_org_anonymous_grpc_CmdGetByNameExtResponse_MsgOnFailure_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_org_anonymous_grpc_CmdGetByNameExtResponse_MsgOnFailure_descriptor,
        new java.lang.String[] { "HasSucceeded", "ErrorType", "ErrorText", });
    org.anonymous.grpc.ErrorTypeOuterClass.getDescriptor();
    org.anonymous.grpc.MetadataOuterClass.getDescriptor();
    org.anonymous.grpc.CmdTypeOuterClass.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
