package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"

	pb "github.com/penk110/micro_in_action/rpc_pb"
)

func main() {
	serviceAddress := "127.0.0.1:8030"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect failed, service: %s err: %s\n", serviceAddress, err.Error())
	}
	defer func(conn *grpc.ClientConn) {
		if err := conn.Close(); err != nil {
			log.Printf("connect close failed, service: %s err: %s\n", serviceAddress, err.Error())
		}
	}(conn)
	stringClient := pb.NewStringServiceClient(conn)

	sendClientRequest(stringClient)

	//sendClientStreamRequest(stringClient)
	//
	//sendClientAndServerStreamRequest(stringClient)
}

func sendClientRequest(client pb.StringServiceClient) {
	var (
		stream pb.StringService_LotsOfServerStreamClient
		ctx    context.Context
		cancel context.CancelFunc
		reply  *pb.StringResponse
		err    error
	)
	fmt.Printf("---------- sendClientRequest start\n")

	stringReq := &pb.StringRequest{A: "A", B: "B"}
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	if stream, err = client.LotsOfServerStream(ctx, stringReq); err != nil {
		log.Printf("failed to call: %v", err)
		return
	}
	for {
		reply, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed to recv: %v", err)
		}
		fmt.Printf("StringService Concat : %s concat %s = %s\n", stringReq.A, stringReq.B, reply.GetResult())
	}

	fmt.Printf("---------- sendClientRequest end\n")
}

func sendClientStreamRequest(client pb.StringServiceClient) {
	fmt.Printf("---------- sendClientStreamRequest\n\n")
	stream, err := client.LotsOfClientStream(context.Background())
	for i := 0; i < 10; i++ {
		if err != nil {
			log.Printf("failed to call: %v", err)
			break
		}

		if err := stream.Send(&pb.StringRequest{A: strconv.Itoa(i), B: strconv.Itoa(i + 1)}); err != nil {
			return
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("failed to recv: %v", err)
	}
	log.Printf("sendClientStreamRequest ret is : %s", reply.Result)
}

func sendClientAndServerStreamRequest(client pb.StringServiceClient) {
	var (
		stream pb.StringService_LotsOfServerAndClientStreamClient
		ctx    context.Context
		cancel context.CancelFunc
		reply  *pb.StringResponse
		err    error
	)
	fmt.Printf("---------- sendClientAndServerStreamRequest\n\n")
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if stream, err = client.LotsOfServerAndClientStream(ctx); err != nil {
		log.Printf("failed to call: %v", err)
		return
	}
	var i int
	for {
		if err = stream.Send(&pb.StringRequest{A: strconv.Itoa(i), B: strconv.Itoa(i + 1)}); err != nil {
			if err == io.EOF {
				log.Printf("[sendClientAndServerStreamRequest] EOF")
				return
			}
			log.Printf("failed to send: %v", err)
			break
		}
		if reply, err = stream.Recv(); err != nil {
			log.Printf("failed to recv: %v", err)
			break
		}
		log.Printf("sendClientAndServerStreamRequest Ret is : %s", reply.Result)
		i++
	}
}
