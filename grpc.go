package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "go-grpc-demo/grpc"

	"google.golang.org/grpc"
)

func NewGrpc(port string) *Grpc {
	var err error
	var g = new(Grpc)
	g.port = port
	g.Client, err = grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Grpc error")
	}
	return g
}

type Grpc struct {
	pb.GrpcSrvServer
	Client *grpc.ClientConn
	port   string
}

func (g *Grpc) Run() {
	port, err := net.Listen("tcp", g.port)
	if err != nil {
		log.Fatal("Listen net error")
	}
	server := grpc.NewServer()
	pb.RegisterGrpcSrvServer(server, g)
	server.Serve(port)
}

func (g *Grpc) GetData(ctx context.Context, req *pb.GetDataReq) (resp *pb.GetDataResp, err error) {
	resp = &pb.GetDataResp{
		Response: fmt.Sprintf("%s: Succeed", req.Request),
	}
	return
}
