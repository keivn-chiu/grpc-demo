package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/kevin-chiu/grpc-demo/api/phone"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("create connection err: %v\n", err)
	}
	defer conn.Close()
	cli := pb.NewPhoneHelperClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	phoneList := []*pb.Phone{
		{Brand: pb.PhoneBrand_Apple, Name: "IPhone 12"},
		{Brand: pb.PhoneBrand_Apple, Name: "IPhone 8 pro"},
		{Brand: pb.PhoneBrand_Apple, Name: "IPhone x"},
		{Brand: pb.PhoneBrand_HuaWei, Name: "P10"},
		{Brand: pb.PhoneBrand_HuaWei, Name: "P20"},
		{Brand: pb.PhoneBrand_Samsung, Name: "Galaxy z"},
		{Brand: pb.PhoneBrand_Samsung, Name: "Galaxy Note"},
	}
	ph, err := cli.Classify(ctx, &pb.Phones{PhonesList: phoneList})
	if err != nil {
		log.Fatalf("classify err: %s\n", err.Error())
	}
	tctx, done := context.WithTimeout(ctx, time.Duration(time.Second*10))
	defer done()
	for {
		select {
		case <-tctx.Done():
			return
		default:
			phones, err := ph.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("receive message from sever error: %v\n", err)
			}
			for _, v := range phones.PhonesList {
				log.Printf("phone brand: %v, name: %s\n", v.Brand.String(), v.Name)
			}
			log.Println()
		}
	}
}
