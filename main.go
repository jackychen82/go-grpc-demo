package main

import (
	"go-grpc-demo/client"
	"go-grpc-demo/server"
	"log"
	"net/http"
)

func main() {
	g := server.NewGrpc(":8081")
	go g.Run()
	r := client.Rest{
		Client: g.Client(),
	}
	http.HandleFunc("/grpc", r.Handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
