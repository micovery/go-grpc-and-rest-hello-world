## Sample Go gRPC / REST Server

This is a sample gRPC / REST server written in Go. 
It uses LabStack's Echo HTTP/2 server, instead of the built-in HTTP/2 server that comes
with Go's gRPC implementation. 

This sample server demonstrates how to run both gRPC / REST. It also shows how to return
trailers headers for both gRPC and REST.


## Pe-Requisites

* Go 1.21 or later
* Protobuf compiler ([protoc](https://github.com/protocolbuffers/protobuf/releases_))
* Go Plugin for protoc
  * go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
* Go gRPC plugin for protoc
  * go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

## How to compile and run


* If you made changes to the [greeter.proto](/pkg/greeter/greeter.proto) file, re-generate gRPC files. 
Make sure you have all `go`, `protoc`, `protoc-gen-go`, and `protoc-gen-go-grpc` in your path.
  ```shell
  go generate ./...
  ```
* Compile the example `hello-world.go`
  ```shell
  go build -o hello-world ./cmd/hello-world.go
  ```
  This will produce a new binary `hello-world` in the current directory


* Finally, you can run the example
  ```shell
  PORT=80 ./hello-world
  ```

## Sample gRPC Requests

* Reflection
  ```shell
  $ grpcurl -v -plaintext  -proto ./pkg/greeter/greeter.proto localhost:80 list
  ```
  ```text
  helloworld.Greeter
  ```

* Unary request
  ```shell
  $ grpcurl -v -plaintext  -proto ./pkg/greeter/greeter.proto localhost:80 helloworld.Greeter.SayHello
  ```
  ```text
  Resolved method descriptor:
  rpc SayHello ( .helloworld.Empty ) returns ( .helloworld.SayHelloRes );
  
  Request metadata to send:
  (empty)
  
  Response headers received:
  content-type: application/grpc
  header1: header-value1
  header2: header-value2
  trailer: Grpc-Status
  trailer: Grpc-Message
  trailer: Grpc-Status-Details-Bin
  
  Response contents:
  {
    "message": "Hello World"
  }
  
  Response trailers received:
  trailer1: trailer-value1
  trailer2: trailer-value2
  ```

## Sample REST Requests

* Using HTTP/1.1 

  ```shell
  $ curl --raw -v --http1.1 http://localhost:80
  ```
  
  ```text
  > GET / HTTP/1.1
  > Host: localhost
  > User-Agent: curl/8.1.2
  > Accept: */*
  > 
  < HTTP/1.1 200 OK
  < Content-Type: application/json; charset=UTF-8
  < Trailer: trailer1, trailer2
  < Date: Mon, 13 Nov 2023 19:01:53 GMT
  < Transfer-Encoding: chunked
  < 
  1a
  {"Message":"Hello World"}
  
  0
  Trailer1: trailer-value1
  Trailer2: trailer-value2
  ```

* Using HTTP/2 (with h2c upgrade)

  ```shell
  $ curl --raw -v --http2 http://localhost:80
  ```
    
  ```text
  > GET / HTTP/1.1
  > Host: localhost
  > User-Agent: curl/8.1.2
  > Accept: */*
  > Connection: Upgrade, HTTP2-Settings
  > Upgrade: h2c
  > HTTP2-Settings: AAMAAABkAAQAoAAAAAIAAAAA
  > 
  < HTTP/1.1 101 Switching Protocols
  < Connection: Upgrade
  < Upgrade: h2c
  * Received 101, Switching to HTTP/2
  < HTTP/2 200 
  < content-type: application/json; charset=UTF-8
  < trailer: trailer1, trailer2
  < date: Mon, 13 Nov 2023 19:03:35 GMT
  < 
  {"Message":"Hello World"}
  < trailer1: trailer-value1
  < trailer2: trailer-value2
  ```
  
* Using HTTP/2 (with prior knowledge)

  ```shell
  $ curl --raw -v --http2 --http2-prior-knowledge http://localhost:80
  ```

  ```text
  * h2 [:method: GET]
  * h2 [:scheme: http]
  * h2 [:authority: localhost]
  * h2 [:path: /]
  * h2 [user-agent: curl/8.1.2]
  * h2 [accept: */*]
  * Using Stream ID: 1 (easy handle 0x14f00a800)
  > GET / HTTP/2
  > Host: localhost
  > User-Agent: curl/8.1.2
  > Accept: */*
  > 
  < HTTP/2 200 
  < content-type: application/json; charset=UTF-8
  < trailer: trailer1, trailer2
  < date: Mon, 13 Nov 2023 19:05:16 GMT
  < 
  {"Message":"Hello World"}
  < trailer1: trailer-value1
  < trailer2: trailer-value2
  ```


### Not Google Product Clause

This is not an officially supported Google product.