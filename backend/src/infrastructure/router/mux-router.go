package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

type muxRouter struct{}

var muxRouterInstance = mux.NewRouter()

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(uri, f).Methods("POST")
	muxRouterInstance.HandleFunc(uri, f).Methods("OPTIONS")
}

func (*muxRouter) SERVE(port string) {
	println("Mux HTTP Sever running on port", port)
	http.ListenAndServe(port, muxRouterInstance)
}
