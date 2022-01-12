package endpoint

import (
	"context"
	"log"
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/penk110/micro_in_action/consul_register_discover/service"
)

type HealthReq struct{}

type HealthResp struct {
	Status bool `json:"status"`
}

// MakeHealthCheckEndpoint 创建健康检查Endpoint
func MakeHealthCheckEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		log.Printf("[MakeHealthCheckEndpoint] time: %d\n", time.Now().Unix())
		status := svc.HealthCheck()
		return HealthResp{
			Status: status,
		}, nil
	}
}
