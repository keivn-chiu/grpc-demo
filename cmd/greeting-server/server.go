package main

import (
	"log"
	"net"

	pb "github.com/kevin-chiu/grpc-demo/api/greeting"
	"github.com/kevin-chiu/grpc-demo/pkg/server/greeting"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen err: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetingServer(s, greeting.NewServer())
	log.Printf("start serve on port: %s\n", port)
	if err := s.Serve(l); err != nil {
		log.Fatalf("serve err: %v\n", err)
	}
}
