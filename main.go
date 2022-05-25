package main

import (
	"context"
	"log"
	"net/http"

	pb "go-grpc-demo/grpc"

	"google.golang.org/grpc"
)

type Rest struct {
	GrpcClient *grpc.ClientConn
}

func (r *Rest) handle(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(r.getData()))
}

func (r *Rest) getData() string {
	client := pb.NewGrpcSrvClient(r.GrpcClient)

	req := &pb.GetDataReq{
		Request: "Get data",
	}

	result, err := client.GetData(context.TODO(), req)
	if err != nil {
		return "ERROR"
	}

	return result.Response
}

func main() {
	g := NewGrpc(":8081")
	go g.Run()
	r := Rest{
		GrpcClient: g.Client,
	}
	http.HandleFunc("/grpc", r.handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
