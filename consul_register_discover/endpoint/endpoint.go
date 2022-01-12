package endpoint

import (
	"github.com/go-kit/kit/endpoint"
)

type DiscoveryEndpoints struct {
	PingPongEndpoint    endpoint.Endpoint
	DiscoveryEndpoint   endpoint.Endpoint
	HealthCheckEndpoint endpoint.Endpoint
}

