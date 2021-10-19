package main

import (
	"log"
	"net"

	greeting_pb "github.com/kevin-chiu/grpc-demo/api/greeting"
	string_pb "github.com/kevin-chiu/grpc-demo/api/string"
	"github.com/kevin-chiu/grpc-demo/pkg/server/greeting"
	stringjoin "github.com/kevin-chiu/grpc-demo/pkg/server/string-join"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("listen err: %v\n", err)
	}
	s := grpc.NewServer()
	greeting_pb.RegisterGreetingServer(s, greeting.NewServer())
	string_pb.RegisterStringJoinServer(s, stringjoin.NewServer())
	log.Println("serve")
	if err := s.Serve(l); err != nil {
		log.Fatalf("serve err: %v\n", err)
	}
}
