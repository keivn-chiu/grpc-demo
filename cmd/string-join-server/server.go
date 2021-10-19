package main

import (
	"log"
	"net"

	pb "github.com/kevin-chiu/grpc-demo/api/string"
	stringjoin "github.com/kevin-chiu/grpc-demo/pkg/server/string-join"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterStringJoinServer(s, stringjoin.NewServer())
	log.Printf("start to listen on port: %v\n", port)
	if err := s.Serve(l); err != nil {
		log.Fatalf("can't serve: %v\n", err)
	}
}
