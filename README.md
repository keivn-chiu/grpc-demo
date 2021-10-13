# GRPC DEMOS

To learn grpc, demos in golang.

## Environment

Mac OS Environment:

1. Install golang -> `brew install go`
2. Install protocol buffer compiler -> `brew install protobuf`
3. Install protocol buffer golang plugin -> `google.golang.org/protobuf/cmd/protoc-gen-go@latest`
4. Add `$GOPATH/bin` into `$PATH` so that compiler can get the golang plugin

## Run Demos

### 1. ProductInfo GRPC

In terminal:

1. To get dependences and create binary file -> make init
2. To generate api file -> make api-gen
3. To build client software -> make product-client
4. To build server software -> make product-server