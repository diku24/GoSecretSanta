package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type muxRouter struct {
}

var muxDispatcher = mux.NewRouter()

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, funobj func(response http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, funobj).Methods("GET")
}
func (*muxRouter) POST(uri string, funobj func(response http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, funobj).Methods("POST")
}
func (*muxRouter) PUT(uri string, funobj func(response http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, funobj).Methods("PUT")
}

func (*muxRouter) SERVE(port string) {
	logrus.Info("MUX HTTP Server is Running on Port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}
