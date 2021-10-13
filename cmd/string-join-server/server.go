package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"strings"
	"time"

	pb "github.com/kevin-chiu/grpc-demo/api/string"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const port = ":50051"

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterStringJoinServer(s, &server{})
	log.Printf("start to listen on port: %v\n", port)
	if err := s.Serve(l); err != nil {
		log.Fatalf("can't serve: %v\n", err)
	}
}

type server struct {
}

func (s *server) Join(js pb.StringJoin_JoinServer) error {
	ctx := context.Background()
	tctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second))
	defer cancel()
	var strs []string
	for {
		select {
		case <-tctx.Done():
			return errors.New("timeout error")
		default:
			str, err := js.Recv()
			if err == io.EOF {
				final := strings.Join(strs, "->")
				return js.SendAndClose(&wrapperspb.StringValue{Value: final}) // here need to return func
			}
			if err != nil {
				return err
			}
			strs = append(strs, str.Value)
		}
	}
}
