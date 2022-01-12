package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	stringService "github.com/penk110/micro_in_action/rpc_grpc_stream/service"
	pb "github.com/penk110/micro_in_action/rpc_pb"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8030")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	ss := new(stringService.StringService)
	pb.RegisterStringServiceServer(grpcServer, ss)
	if err := grpcServer.Serve(lis); err != nil {
		lis.Close()
		panic(err)
	}
}
