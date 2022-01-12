package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/penk110/micro_in_action/consul_register_discover/service"
)

type DiscoveryReq struct {
	ServiceName string
}

type DiscoveryResp struct {
	Instances []interface{} `json:"instances"`
	Error     string        `json:"error"`
}

func MakeDiscoveryEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var (
			req DiscoveryReq
		)
		req = request.(DiscoveryReq)
		instances, err := svc.DiscoveryService(ctx, req.ServiceName)
		var errString = ""
		if err != nil {
			errString = err.Error()
		}
		return &DiscoveryResp{
			Instances: instances,
			Error:     errString,
		}, nil
	}
}
