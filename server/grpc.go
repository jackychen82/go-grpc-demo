package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"go-grpc-demo/proto"

	"google.golang.org/grpc"
)

func NewGrpc(port string) *Grpc {
	var err error
	var g = new(Grpc)
	g.port = port
	g.client, err = grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Grpc error")
	}
	return g
}

type Grpc struct {
	client *grpc.ClientConn
	port   string

	proto.GreeterServer
}

func (g *Grpc) Client() *grpc.ClientConn {
	return g.client
}

func (g *Grpc) Run() {
	port, err := net.Listen("tcp", g.port)
	if err != nil {
		log.Fatal("Listen net error")
	}
	serv := grpc.NewServer()
	proto.RegisterGreeterServer(serv, g)
	serv.Serve(port)
}

func (g *Grpc) SayHello(ctx context.Context, req *proto.HelloRequest) (resp *proto.HelloReply, err error) {
	resp = &proto.HelloReply{
		Message: fmt.Sprintf("Hello! %s!", req.Name),
	}
	return
}
