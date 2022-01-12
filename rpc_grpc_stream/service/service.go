package service

import (
	"context"
	"io"
	"log"
	"strings"

	"github.com/penk110/micro_in_action/rpc_pb"
)

const (
	StrMaxSize = 1024
)

type StringService struct {
}

/*
	Concat(context.Context, *StringRequest) (*StringResponse, error)
发送：
	LotsOfServerStream(*StringRequest, StringService_LotsOfServerStreamServer) error
接收：
	LotsOfClientStream(StringService_LotsOfClientStreamServer) error
接收和发送：
	LotsOfServerAndClientStream(StringService_LotsOfServerAndClientStreamServer) error

*/

func (ss *StringService) Concat(ctx context.Context, req *stream_pb.StringRequest) (*stream_pb.StringResponse, error) {
	if len(req.A)+len(req.B) > StrMaxSize {
		response := stream_pb.StringResponse{Result: ""}
		return &response, nil
	}
	response := stream_pb.StringResponse{Result: req.A + req.B}
	return &response, nil
}

func (ss *StringService) LotsOfServerStream(req *stream_pb.StringRequest, qs stream_pb.StringService_LotsOfServerStreamServer) error {
	resp := stream_pb.StringResponse{Result: req.A + req.B}
	for i := 0; i < 10; i++ {
		if err := qs.Send(&resp); err != nil {
			log.Printf("[LotsOfServerStream] failed to send, err: %s", err.Error())
			return err
		}
	}
	return nil
}

func (ss *StringService) LotsOfClientStream(qs stream_pb.StringService_LotsOfClientStreamServer) error {
	var params []string
	for {
		in, err := qs.Recv()
		if err == io.EOF {
			if err := qs.SendAndClose(&stream_pb.StringResponse{Result: strings.Join(params, "")}); err != nil {
				log.Printf("[LotsOfClientStream.SendAndClose] failed to close, err: %s", err.Error())
				return nil
			}
			return nil
		}
		if err != nil {
			log.Printf("[LotsOfClientStream] failed to recv, err: %s", err.Error())
			return err
		}
		params = append(params, in.A, in.B)
		log.Printf("[LotsOfClientStream]recv, A: %s B: %s", in.A, in.B)
	}
}

func (ss *StringService) LotsOfServerAndClientStream(qs stream_pb.StringService_LotsOfServerAndClientStreamServer) error {
	for {
		in, err := qs.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("[LotsOfServerAndClientStream] failed to recv, err: %s", err.Error())
			return err
		}
		if err := qs.Send(&stream_pb.StringResponse{Result: in.A + in.B}); err != nil {
			log.Printf("[LotsOfServerAndClientStream] failed to send, err: %s", err.Error())
			return err
		}
	}
}
