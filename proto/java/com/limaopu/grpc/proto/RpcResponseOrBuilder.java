// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: service.proto

package com.limaopu.grpc.proto;

public interface RpcResponseOrBuilder extends
    // @@protoc_insertion_point(interface_extends:proto.RpcResponse)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>int32 code = 1;</code>
   * @return The code.
   */
  int getCode();

  /**
   * <code>string message = 2;</code>
   * @return The message.
   */
  java.lang.String getMessage();
  /**
   * <code>string message = 2;</code>
   * @return The bytes for message.
   */
  com.google.protobuf.ByteString
      getMessageBytes();

  /**
   * <code>bytes data = 3;</code>
   * @return The data.
   */
  com.google.protobuf.ByteString getData();
}
