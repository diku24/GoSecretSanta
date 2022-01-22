package api

import "net/http"

type Router interface {
	GET(uri string, funobj func(response http.ResponseWriter, request *http.Request))
	POST(uri string, funobj func(response http.ResponseWriter, request *http.Request))
	PUT(uri string, funobj func(response http.ResponseWriter, request *http.Request))
	SERVE(port string)
}
