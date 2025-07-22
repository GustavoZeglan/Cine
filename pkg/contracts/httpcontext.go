package contracts

import (
	"context"
	"net/http"
)

type HttpContext struct {
	Writer  http.ResponseWriter
	Request http.Request
	Context context.Context
	Params  map[string]string
}
