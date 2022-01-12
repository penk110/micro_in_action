package service

import (
	"context"
)

// Service impl
type Service interface {
	// HealthCheck check service health status
	HealthCheck() bool
	// PingPong ping pong service
	PingPong() string
	// DiscoveryService discovery service from consul by serviceName
	DiscoveryService(ctx context.Context, serviceName string) ([]interface{}, error)
}
