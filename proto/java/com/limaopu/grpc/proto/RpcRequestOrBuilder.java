// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: service.proto

package com.limaopu.grpc.proto;

public interface RpcRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:proto.RpcRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string route = 1;</code>
   * @return The route.
   */
  java.lang.String getRoute();
  /**
   * <code>string route = 1;</code>
   * @return The bytes for route.
   */
  com.google.protobuf.ByteString
      getRouteBytes();

  /**
   * <code>string token = 2;</code>
   * @return The token.
   */
  java.lang.String getToken();
  /**
   * <code>string token = 2;</code>
   * @return The bytes for token.
   */
  com.google.protobuf.ByteString
      getTokenBytes();

  /**
   * <code>bytes data = 3;</code>
   * @return The data.
   */
  com.google.protobuf.ByteString getData();
}
