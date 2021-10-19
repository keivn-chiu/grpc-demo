package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kevin-chiu/grpc-demo/api/string"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial err: %v\n", err)
	}
	sjc := pb.NewStringJoinClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sj, err := sjc.Join(ctx)
	if err != nil {
		log.Fatalf("join err: %v\n", err)
	}
	time.Sleep(time.Second)
	cancel()

	_, err = sj.CloseAndRecv()
	log.Printf("err: %v", err)
}
