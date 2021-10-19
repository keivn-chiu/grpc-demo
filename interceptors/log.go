package interceptors

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func LogUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	log.Printf("start\n")
	resp, err = handler(ctx, req)
	log.Printf("cost: %d\n", time.Since(start).Nanoseconds())
	return
}

type wrappedStream struct {
	grpc.ServerStream
}

func newWrappedStream(ss grpc.ServerStream) *wrappedStream {
	return &wrappedStream{ss}
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	println("===== Server Interceptor Wrapper Recv =====")
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	println("===== Server Interceptor Wrapper Send =====")
	return w.ServerStream.SendMsg(m)
}

func LogStreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("request get")
	start := time.Now()
	err := handler(srv, newWrappedStream(ss))
	log.Printf("cost: %d\n", time.Since(start).Nanoseconds())
	return err
}