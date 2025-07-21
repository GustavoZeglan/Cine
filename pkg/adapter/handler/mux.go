package adapter

import (
	"net/http"

	"github.com/GustavoZeglan/Cine/pkg/contracts"
	"github.com/gorilla/mux"
)

type MuxAdapter struct {
	Router *mux.Router
}

func NewMuxAdapter() *MuxAdapter {
	return &MuxAdapter{
		Router: mux.NewRouter(),
	}
}

var _ contracts.HttpMethod = (*MuxAdapter)(nil)

func wrapMux(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
	}
}

func (m *MuxAdapter) Delete(path string, handler http.HandlerFunc) {
	m.Router.HandleFunc(path, wrapMux(handler)).Methods(http.MethodDelete)
}

func (m *MuxAdapter) Get(path string, handler http.HandlerFunc) {
	m.Router.HandleFunc(path, wrapMux(handler)).Methods(http.MethodGet)
}

func (m *MuxAdapter) Patch(path string, handler http.HandlerFunc) {
	m.Router.HandleFunc(path, wrapMux(handler)).Methods(http.MethodPatch)
}

func (m *MuxAdapter) Post(path string, handler http.HandlerFunc) {
	m.Router.HandleFunc(path, wrapMux(handler)).Methods(http.MethodPost)
}

func (m *MuxAdapter) Put(path string, handler http.HandlerFunc) {
	m.Router.HandleFunc(path, wrapMux(handler)).Methods(http.MethodPut)
}
