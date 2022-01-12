package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/penk110/micro_in_action/rpc_native/server"
)

func main() {
	stringServer := new(server.StringServer)
	registerErr := rpc.Register(stringServer)
	if registerErr != nil {
		panic(registerErr)
	}

	rpc.HandleHTTP()
	l, err := net.Listen("tcp", "127.0.0.1:8030")
	if err != nil {
		log.Fatalf("listen failed, err: %s", err)
	}

	if err := http.Serve(l, nil); err != nil {
		log.Fatalf("start server failed, err: %s", err.Error())
		return
	}
}
