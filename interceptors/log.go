package interceptors

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func LogUnaryServerInterceptorfunc(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	log.Printf("start\n")
	resp, err = handler(ctx, req)
	log.Printf("cost: %d\n", time.Since(start).Nanoseconds())
	return
}
