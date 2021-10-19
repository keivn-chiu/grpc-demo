package main

import (
	"errors"
	"log"
	"net"

	pb "github.com/kevin-chiu/grpc-demo/api/phone"
	"github.com/kevin-chiu/grpc-demo/interceptors"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen, %v\n", err)
	}
	s := grpc.NewServer(grpc.StreamInterceptor(interceptors.LogStreamServerInterceptor))
	pb.RegisterPhoneHelperServer(s, &server{})
	log.Printf("start to listen port: %s\n", port)
	if err := s.Serve(l); err != nil {
		log.Fatalf("listen failed: %v\n", err)
	}
}

type server struct {
}

func (s *server) Classify(ps *pb.Phones, cs pb.PhoneHelper_ClassifyServer) error {
	list := ps.PhonesList
	if len(list) == 0 {
		return errors.New("invalid phone list")
	}
	m := make(map[pb.PhoneBrand]*pb.Phones)
	for _, v := range list {
		p, ok := m[v.Brand]
		if !ok {
			m[v.Brand] = &pb.Phones{PhonesList: []*pb.Phone{v}}
			continue
		}
		p.PhonesList = append(p.PhonesList, v)
		m[v.Brand] = p
	}
	for _, v := range m {
		err := cs.Send(v)
		if err != nil {
			return err
		}
	}
	return nil
}
