package main

import (
	"context"
	"log"

	pb "github.com/kevin-chiu/grpc-demo/api/string"
	"github.com/kevin-chiu/grpc-demo/interceptors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithStreamInterceptor(interceptors.LogStreamClientInterceptor))
	if err != nil {
		log.Fatalf("dial failed: %v\n", err)
	}
	cli := pb.NewStringJoinClient(conn)

	scli, err := cli.Join(context.Background())
	if err != nil {
		log.Fatalf("call func err: %v\n", err)
	}
	err = scli.Send(wrapperspb.String("Hello"))
	if err != nil {
		log.Fatalf("send failed: %v\n", err)
	}
	err = scli.Send(wrapperspb.String("World"))
	if err != nil {
		log.Fatalf("send failed: %v\n", err)
	}
	err = scli.Send(wrapperspb.String("!"))
	if err != nil {
		log.Fatalf("send failed: %v\n", err)
	}
	ret, err := scli.CloseAndRecv()
	if err != nil {
		log.Fatalf("try rec failed: %v\n", err)
	}
	log.Printf("server returns: %s\n", ret.Value)

}
