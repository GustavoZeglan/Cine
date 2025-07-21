package contracts

import "net/http"

type HttpMethod interface {
	Get(path string, handler http.HandlerFunc)
	Post(path string, handler http.HandlerFunc)
	Put(path string, handler http.HandlerFunc)
	Delete(path string, handler http.HandlerFunc)
	Patch(path string, handler http.HandlerFunc)
}
