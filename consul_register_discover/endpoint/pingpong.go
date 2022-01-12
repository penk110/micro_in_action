package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/penk110/micro_in_action/consul_register_discover/service"
)

type PingResp struct {
	Message string `json:"message"`
	Result  string `json:"result"`
}

func MakePingPongEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var pingResp = &PingResp{
			Message: svc.PingPong(),
			Result:  "success",
		}
		return pingResp, nil
	}
}
