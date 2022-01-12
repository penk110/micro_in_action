package string_service

import (
	"context"
)

type Service interface {
	Concat(ctx context.Context, a, b string) (string, error)
	Diff(ctx context.Context, a, b string) (string, error)
}

type ServiceMiddleware func(Service) Service
