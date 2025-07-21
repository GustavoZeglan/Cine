package adapter

import (
	"net/http"

	"github.com/GustavoZeglan/Cine/pkg/contracts"
	"github.com/gin-gonic/gin"
)

type GinAdapter struct {
	Router *gin.Engine
}

func NewGinAdapter() *GinAdapter {
	return &GinAdapter{
		Router: gin.Default(),
	}
}

var _ contracts.HttpMethod = (*GinAdapter)(nil)

func wrapGin(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c.Writer, c.Request)
	}
}

func (g *GinAdapter) Delete(path string, handler http.HandlerFunc) {
	g.Router.DELETE(path, wrapGin(handler))
}

func (g *GinAdapter) Get(path string, handler http.HandlerFunc) {
	g.Router.GET(path, wrapGin(handler))
}

func (g *GinAdapter) Patch(path string, handler http.HandlerFunc) {
	g.Router.PATCH(path, wrapGin(handler))
}

func (g *GinAdapter) Post(path string, handler http.HandlerFunc) {
	g.Router.POST(path, wrapGin(handler))
}

func (g *GinAdapter) Put(path string, handler http.HandlerFunc) {
	g.Router.PUT(path, wrapGin(handler))
}
