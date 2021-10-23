<?php
// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
namespace Proto;

/**
 */
class MyServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Proto\RpcRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function RpcExecute(\Proto\RpcRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/proto.MyService/RpcExecute',
        $argument,
        ['\Proto\RpcResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \Proto\RpcJsonRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function RpcJsonExecute(\Proto\RpcJsonRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/proto.MyService/RpcJsonExecute',
        $argument,
        ['\Proto\RpcJsonResponse', 'decode'],
        $metadata, $options);
    }

}
