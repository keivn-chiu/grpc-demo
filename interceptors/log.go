package interceptors

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func LogUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	log.Printf("start\n")
	resp, err = handler(ctx, req)
	log.Printf("cost: %d\n", time.Since(start).Nanoseconds())
	return
}

type wrappedServerStream struct {
	grpc.ServerStream
}

func newWrappedStream(ss grpc.ServerStream) *wrappedServerStream {
	return &wrappedServerStream{ss}
}

func (w *wrappedServerStream) RecvMsg(m interface{}) error {
	println("===== Server Interceptor Wrapper Recv =====")
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedServerStream) SendMsg(m interface{}) error {
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

func LogUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	log.Println("unary client -> before send requests")
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("unary client -> after send requests cost: %d, reply: %v\n", time.Since(start).Nanoseconds(), reply)
	return err
}

type wrappedClientStream struct {
	grpc.ClientStream
}

func newWrappedClientStream(cs grpc.ClientStream) *wrappedClientStream {
	return &wrappedClientStream{cs}
}

func (w *wrappedClientStream) RecvMsg(m interface{}) error {
	println("===== Client Interceptor Wrapper Recv =====")
	return w.ClientStream.RecvMsg(m)
}
func (w *wrappedClientStream) SendMsg(m interface{}) error {
	println("===== Client Interceptor Wrapper Send =====")
	return w.ClientStream.SendMsg(m)
}

func LogStreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Printf("before do: %s\n", method)
	// Streamer is called by StreamClientInterceptor to create a ClientStream.
	cs, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrappedClientStream(cs), err
}

func MetadataUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md := metadata.Pairs("name", "kevin", "age", "30")
	c := metadata.NewOutgoingContext(ctx, md)
	err := invoker(c, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}

	return nil
}

func MetadataUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	m, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Printf("server get md: %v\n", m)
	}
	m.Delete("age")
	m.Append("height", "178")
	grpc.SendHeader(ctx, m)
	md := metadata.New(map[string]string{"end": "true"})
	grpc.SetTrailer(ctx, md)
	resp, err = handler(ctx, req)
	return
}
