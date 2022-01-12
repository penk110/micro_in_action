package transport

import (
	"context"
	"net/http"

	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	endpoints "github.com/penk110/micro_in_action/consul_register_discover/endpoint"
)

func regHealthRouter(router *mux.Router, endpoints endpoints.DiscoveryEndpoints, options []kitHttp.ServerOption) {
	router.Methods("GET").Path("/health").Handler(kitHttp.NewServer(
		endpoints.HealthCheckEndpoint,
		decodeHealthCheckRequest,
		encodeJsonResponse,
		options...,
	))
}

func decodeHealthCheckRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return endpoints.HealthReq{}, nil
}
