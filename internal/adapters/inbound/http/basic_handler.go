package http

import (
	"net/http"

	"github.com/pgnahuel/go6arc/internal/ports"
)

type BasicHandler struct {
	service ports.BasicService
}

func NewBasicHandler(service ports.BasicService) ports.BasicHttpHandler {
	return &BasicHandler{
		service: service,
	}
}

func (bh BasicHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}

func (bh BasicHandler) Read(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}

func (bh BasicHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}

func (bh BasicHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}

func (bh BasicHandler) List(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}
