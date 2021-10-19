package greeting

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/kevin-chiu/grpc-demo/api/greeting"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type server struct {
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHiToSomeone(g pb.Greeting_SayHiToSomeoneServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)
	// send func
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return errors.New("cancel")
			default:
				err := g.Send(&wrapperspb.StringValue{Value: "Heart Beat ..."})
				if err != nil {
					cancel()
					return err
				}
				time.Sleep(time.Second * 3)
			}
		}
	})
	// receive func
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return errors.New("cancel")
			default:
				sv, err := g.Recv()
				if err == io.EOF {
					log.Println("receive eof")
					cancel()
					return nil
				}
				if err != nil {
					cancel()
					return err
				}
				err = g.Send(&wrapperspb.StringValue{Value: fmt.Sprintf("Hi %s", sv.Value)})
				if err != nil {
					cancel()
					return err
				}
			}
		}
	})

	return eg.Wait()
}
