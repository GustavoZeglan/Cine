package adapter

import (
	"github.com/GustavoZeglan/Cine/pkg/contracts"
	"github.com/gin-gonic/gin"
)

func Handler(h func(*contracts.HttpContext)) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := make(map[string]string)
		for _, p := range c.Params {
			params[p.Key] = p.Value
		}

		httpCtx := &contracts.HttpContext{
			Writer:  c.Writer,
			Request: *c.Request,
			Context: c.Request.Context(),
			Params:  params,
		}

		h(httpCtx)
	}
}
