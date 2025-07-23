package adapter

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/GustavoZeglan/Cine/pkg/helper"
	"github.com/gin-gonic/gin"
)

type HttpContext struct {
	Writer  http.ResponseWriter
	Request http.Request
	Context context.Context
	Params  map[string]string
}

func NewHttpContext(c *gin.Context) *HttpContext {
	return &HttpContext{
		Writer:  c.Writer,
		Request: *c.Request,
		Context: c.Request.Context(),
		Params:  make(map[string]string),
	}
}

type APIFunc func(*HttpContext) error

func Handler(h APIFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := make(map[string]string)
		for _, p := range c.Params {
			params[p.Key] = p.Value
		}

		httpCtx := NewHttpContext(c)

		if err := h(httpCtx); err != nil {
			if apiErr, ok := err.(*helper.APIError); ok {
				helper.WriteJSON(c.Writer, apiErr.StatusCode, apiErr.Message)
				return
			} else {
				errResp := map[string]any{
					"statusCode": http.StatusInternalServerError,
					"msg":        "internal server error",
				}
				helper.WriteJSON(c.Writer, http.StatusInternalServerError, errResp)
			}
			slog.Error("HTTP API Error", "err", err.Error(), "path", c.Request.URL.Path)
		}
	}
}
