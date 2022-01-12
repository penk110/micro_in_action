package transport

import (
	"context"
	"net/http"

	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	endpoints "github.com/penk110/micro_in_action/consul_register_discover/endpoint"
)

func regPingPongRouter(router *mux.Router, endpoints endpoints.DiscoveryEndpoints, options []kitHttp.ServerOption) {
	router.Methods("GET").Path("/ping").Handler(kitHttp.NewServer(
		endpoints.PingPongEndpoint,
		decodePingPongRequest,
		encodeJsonResponse,
		options...,
	))
}

func decodePingPongRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return endpoints.PingResp{}, nil
}
