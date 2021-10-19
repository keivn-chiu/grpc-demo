package stringjoin

import (
	"context"
	"errors"
	"io"
	"log"
	"strings"
	"time"

	pb "github.com/kevin-chiu/grpc-demo/api/string"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type server struct {
}

func NewServer() *server {
	return &server{}
}

func (s *server) Join(js pb.StringJoin_JoinServer) error {
	ctx := context.Background()
	tctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second))
	defer cancel()
	var strs []string
	for {
		select {
		case <-tctx.Done():
			return errors.New("timeout error")
		default:
			str, err := js.Recv()
			if err == io.EOF {
				// simulate handling delay
				time.Sleep(time.Second)
				final := strings.Join(strs, "->")
				return js.SendAndClose(&wrapperspb.StringValue{Value: final}) // here need to return func
			}
			if err != nil {
				// here can get client context error
				log.Println(errors.Is(js.Context().Err(), context.Canceled))
				return err
			}
			strs = append(strs, str.Value)
		}
	}
}
