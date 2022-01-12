package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/penk110/micro_in_action/rpc_pb"
	rpcStringServer "github.com/penk110/micro_in_action/rpc_string_server"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8030")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	stringService := new(rpcStringServer.StringService)
	pb.RegisterStringServiceServer(grpcServer, stringService)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
