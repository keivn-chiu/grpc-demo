package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kevin-chiu/grpc-demo/api/greeting"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial err: %s\n", err.Error())
	}
	cli := pb.NewGreetingClient(conn)

	gctx := context.Background()
	ctx, cancel := context.WithCancel(gctx)
	defer cancel()
	g, err := cli.SayHiToSomeone(ctx)
	if err != nil {
		log.Fatalf("call err: %s\n", err.Error())
	}
	// send func
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := g.Send(&wrapperspb.StringValue{Value: "Peter"})
				if err != nil {
					log.Fatalf("send err: %v\n", err)
					cancel()
					return
				}
				time.Sleep(time.Second)
			}
		}
	}()
	// receive func
	for {
		select {
		case <-ctx.Done():
			return
		default:
			sv, err := g.Recv()
			if err != nil {
				log.Fatalf("receive err: %v\n", err)
				cancel()
				return
			}
			log.Printf("server -> %s\n", sv.Value)
		}
	}
}
