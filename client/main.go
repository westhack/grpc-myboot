/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"github.com/golang/protobuf/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "grpc-myboot/proto"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMyServiceClient(conn)

	user := &pb.User{
		Id: 2,
		Username: "admin2",
	}

	userBytes, _ := proto.Marshal(user)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RpcExecute(ctx, &pb.RpcRequest{Route: "/getUser/1/admin", Data: userBytes, Token: "myBoot"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	} else if r.Code == 401 {
		log.Fatalf("could not greet: %v", r)
	} else {
		var u = &pb.User{}
		proto.Unmarshal(r.Data, u)
		//var u = &system.SysUser{}
		//s, _ := u.Decode(r.Data)
		log.Printf("Greeting: id %s", u.Id)
		log.Printf("Greeting: name  %s", u.Username)
		log.Printf("Greeting: %v", u)
	}
}
