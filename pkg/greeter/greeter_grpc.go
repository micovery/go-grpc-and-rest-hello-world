// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package greeter

//go:generate protoc  --proto_path=. --go_out=./generated/pb --go-grpc_out=./generated/pb  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative greeter.proto

import (
	"context"
	"github.com/micovery/go-grpc-and-rest-hello-world/pkg/greeter/generated/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GRPCServer struct {
	pb.UnimplementedGreeterServer
}

func (s *GRPCServer) SayHello(ctx context.Context, _ *pb.Empty) (*pb.SayHelloRes, error) {
	defer func() {
		trailer := metadata.Pairs("trailer1", "trailer-value1", "trailer2", "trailer-value2")
		grpc.SetTrailer(ctx, trailer)
	}()

	headers := metadata.Pairs("header1", "header-value1", "header2", "header-value2")
	grpc.SetHeader(ctx, headers)

	return &pb.SayHelloRes{
		Message: "Hello World",
	}, nil
}
