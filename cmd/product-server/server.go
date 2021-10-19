package main

import (
	"context"
	"errors"
	"log"
	"net"
	"strconv"

	pb "github.com/kevin-chiu/grpc-demo/api/product"
	"github.com/kevin-chiu/grpc-demo/interceptors"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %s\n", err.Error())
		return
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(interceptors.LogUnaryServerInterceptorfunc))
	pb.RegisterProductInfoServer(s, &server{})
	log.Printf("start grpc listener on port %s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}

type server struct {
	products []*pb.Product
}

func (s *server) AddProduct(ctx context.Context, product *pb.Product) (*pb.ProductId, error) {
	s.products = append(s.products, product)
	return &pb.ProductId{Value: strconv.Itoa(len(s.products) - 1)}, nil
}
func (s *server) GetProduct(ctx context.Context, id *pb.ProductId) (*pb.Product, error) {
	pId, err := strconv.Atoi(id.Value)
	if err != nil {
		return nil, err
	}
	if pId > len(s.products)-1 {
		return nil, errors.New("invalid index")
	}
	return s.products[pId], nil
}
