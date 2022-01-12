package transport

import (
	"github.com/gorilla/mux"
)

var (
	muxRouter *mux.Router
)

func GetMuxRouter() *mux.Router {
	return muxRouter
}

func init() {
	muxRouter = mux.NewRouter()

	// TODO: do something init
}
