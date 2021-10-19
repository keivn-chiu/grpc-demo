package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kevin-chiu/grpc-demo/api/string"
	"github.com/kevin-chiu/grpc-demo/interceptors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithStreamInterceptor(interceptors.LogStreamClientInterceptor),
	)
	if err != nil {
		log.Fatalf("dial failed: %v\n", err)
	}
	defer conn.Close()
	cli := pb.NewStringJoinClient(conn)

	// add deadline context
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()
	scli, err := cli.Join(ctx)
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
