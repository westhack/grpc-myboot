<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: service.proto

namespace GPBMetadata;

class Service
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�
service.protoproto"8

RpcRequest
route (	
token (	
data (":
RpcResponse
code (
message (	
data ("-
RpcJsonRequest
route (	
data (	">
RpcJsonResponse
code (
message (	
data (	2�
	MyService5

RpcExecute.proto.RpcRequest.proto.RpcResponse" A
RpcJsonExecute.proto.RpcJsonRequest.proto.RpcJsonResponse" B=
com.limaopu.grpc.protoBMyServiceProtoPZgrpc-myboot/protobproto3'
        , true);

        static::$is_initialized = true;
    }
}

