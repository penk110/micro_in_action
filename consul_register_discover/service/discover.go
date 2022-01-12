package service

import (
	"context"

	"github.com/penk110/micro_in_action/consul_register_discover/config"
	"github.com/penk110/micro_in_action/consul_register_discover/discover"
)

func NewDiscoveryService(discoveryClient discover.DiscoveryClient) Service {
	return &DiscoveryService{
		discoveryClient: discoveryClient,
	}
}

type DiscoveryService struct {
	discoveryClient discover.DiscoveryClient
}

func (ds *DiscoveryService) PingPong() string {
	return "pong"
}

func (ds *DiscoveryService) DiscoveryService(ctx context.Context, serviceName string) ([]interface{}, error) {
	instances := ds.discoveryClient.DiscoverServices(serviceName, config.Logger)
	if instances == nil || len(instances) == 0 {
		return nil, ErrNotServiceInstances
	}
	return instances, nil
}

// HealthCheck implement Service method
func (ds *DiscoveryService) HealthCheck() bool {
	return true
}
