package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kevin-chiu/grpc-demo/api/product"
	"github.com/kevin-chiu/grpc-demo/interceptors"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptors.LogUnaryClientInterceptor))
	if err != nil {
		log.Fatalf("can't connect: %v\n", err)
	}
	defer conn.Close()
	cli := pb.NewProductInfoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second))
	defer cancel()
	// try to add product
	fakeIPhone := &pb.Product{Name: "IPhone 11", Id: "11", Description: "Fake IPhone 11"}
	id, err := cli.AddProduct(ctx, fakeIPhone)
	if err != nil {
		log.Fatalf("can't add product %v\n", err)
	}
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
