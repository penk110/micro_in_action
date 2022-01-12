package main

import (
	"context"
	"google.golang.org/grpc"
	"log"

	pb "github.com/penk110/micro_in_action/rpc_pb"
)

func main() {
	serviceAddress := "127.0.0.1:8030"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer func(conn *grpc.ClientConn) {
		if err := conn.Close(); err != nil {
			log.Printf("close conn failed, err: %s", err.Error())
		}
	}(conn)
	bookClient := pb.NewStringServiceClient(conn)
	stringReq := &pb.StringRequest{A: "A", B: "B"}
	reply, err := bookClient.Concat(context.Background(), stringReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("StringService Concat : %s concat %s = %s", stringReq.A, stringReq.B, reply.Result)
}
