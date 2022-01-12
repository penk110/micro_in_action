package transport

import (
	"context"
	"net/http"

	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	endpoints "github.com/penk110/micro_in_action/consul_register_discover/endpoint"
)

func regDiscoverRouter(router *mux.Router, endpoints endpoints.DiscoveryEndpoints, options []kitHttp.ServerOption) {
	router.Methods("GET").Path("/discovery").Handler(kitHttp.NewServer(
		endpoints.DiscoveryEndpoint,
		decodeDiscoveryRequest,
		encodeJsonResponse,
		options...,
	))
}

func decodeDiscoveryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	serviceName := r.URL.Query().Get("serviceName")
	if serviceName == "" {
		return nil, ErrorBadRequest
	}
	return endpoints.DiscoveryReq{
		ServiceName: serviceName,
	}, nil
}
