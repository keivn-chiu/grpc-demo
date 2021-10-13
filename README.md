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

Request <-> Response

In terminal:

1. To get dependences and create binary file -> make init
2. To generate api file -> make api-gen
3. To build client software -> make product-client
4. To build server software -> make product-server

### 2. PhoneClassify GRPC

Request <-> Response stream

In terminal:

1. To get dependences and create binary file -> make init
2. To generate api file -> make api-gen
3. To build client software -> make phone-classify-client
4. To build server software -> make phone-classify-server

### 3. StringJoin GRPC

Request stream <-> Response

In terminal:

1. To get dependences and create binary file -> make init
2. To generate api file -> make api-gen
3. To build client software -> make string-join-client
4. To build server software -> make string-join-server

### 4. Greeting FRPC

Request stream <-> Response stream

In terminal:

1. To get dependences and create binary file -> make init
2. To generate api file -> make api-gen
3. To build client software -> make greeting-client
4. To build server software -> make greeting-server