package client

import (
	"context"
	"go-grpc-demo/proto"
	"net/http"

	"google.golang.org/grpc"
)

type Rest struct {
	Client *grpc.ClientConn
}

func (r *Rest) Handle(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(r.sayHello()))
}

func (r *Rest) sayHello() string {
	client := proto.NewGreeterClient(r.Client)

	req := &proto.HelloRequest{
		Name: "Jack",
	}

	result, err := client.SayHello(context.TODO(), req)
	if err != nil {
		return err.Error()
	}

	return result.Message
}
