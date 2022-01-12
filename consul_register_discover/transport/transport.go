package transport

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	endpoints "github.com/penk110/micro_in_action/consul_register_discover/endpoint"
	"net/http"
)

var (
	router          *mux.Router
	options         []kitHttp.ServerOption
	ErrorBadRequest = errors.New("invalid request parameter")
)

// MakeHttpHandler make http handler use mux
func MakeHttpHandler(ctx context.Context, endpoints endpoints.DiscoveryEndpoints, logger log.Logger) http.Handler {
	router = mux.NewRouter()
	options = []kitHttp.ServerOption{
		kitHttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kitHttp.ServerErrorEncoder(encodeError),
	}

	regPingPongRouter(router, endpoints, options)
	regHealthRouter(router, endpoints, options)
	regDiscoverRouter(router, endpoints, options)

	return router
}

func encodeJsonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	
	// TODO: error
	if err = json.NewEncoder(w).Encode(map[string]interface{}{
		"message": err.Error(),
		"result":  "InternalError",
	}); err != nil {
		return
	}
}
