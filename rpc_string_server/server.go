package string_service

import (
	"context"
	"strings"

	pb "github.com/penk110/micro_in_action/rpc_pb"
)

const (
	StrMaxSize = 1024
)

type StringService struct{}

func (s *StringService) LotsOfServerStream(request *pb.StringRequest, server pb.StringService_LotsOfServerStreamServer) error {
	panic("implement me")
}

func (s *StringService) LotsOfClientStream(server pb.StringService_LotsOfClientStreamServer) error {
	panic("implement me")
}

func (s *StringService) LotsOfServerAndClientStream(server pb.StringService_LotsOfServerAndClientStreamServer) error {
	panic("implement me")
}

func (s *StringService) Concat(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	if len(req.A)+len(req.B) > StrMaxSize {
		response := pb.StringResponse{Result: ""}
		return &response, nil
	}
	response := pb.StringResponse{Result: req.A + req.B}
	return &response, nil
}

func (s *StringService) Diff(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	if len(req.A) < 1 || len(req.B) < 1 {
		response := pb.StringResponse{Result: ""}
		return &response, nil
	}
	res := ""
	if len(req.A) >= len(req.B) {
		for _, char := range req.B {
			if strings.Contains(req.A, string(char)) {
				res = res + string(char)
			}
		}
	} else {
		for _, char := range req.A {
			if strings.Contains(req.B, string(char)) {
				res = res + string(char)
			}
		}
	}
	response := pb.StringResponse{Result: res}
	return &response, nil
}
