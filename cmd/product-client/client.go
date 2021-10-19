package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kevin-chiu/grpc-demo/api/product"
	"github.com/kevin-chiu/grpc-demo/interceptors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptors.LogUnaryClientInterceptor),
		grpc.WithUnaryInterceptor(interceptors.MetadataUnaryClientInterceptor),
	)
	if err != nil {
		log.Fatalf("can't connect: %v\n", err)
	}
	defer conn.Close()
	cli := pb.NewProductInfoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second))
	defer cancel()
	// try to add product
	fakeIPhone := &pb.Product{Name: "IPhone 11", Id: "11", Description: "Fake IPhone 11"}
	// try to get metadata as well
	var header, trailer metadata.MD
	id, err := cli.AddProduct(ctx, fakeIPhone, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatalf("can't add product %v\n", err)
	}
	log.Printf("header: %v\n", header)
	log.Printf("trailer: %v\n", trailer)
	log.Printf("Add Product Successful -> ID = %v\n", id.Value)

	// try to get product just added
	product, err := cli.GetProduct(ctx, &pb.ProductId{Value: "0"})
	if err != nil {
		log.Fatalf("can't get %s product %v\n", "0", err)
	}
	log.Printf("Get product: %v\n", product)

	product, err = cli.GetProduct(ctx, &pb.ProductId{Value: "1"})
	if err != nil {
		log.Fatalf("can't get %s product %v\n", "0", err)
	}
	log.Printf("Get product: %v\n", product)
}
