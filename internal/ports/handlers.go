package ports

import "net/http"

type BasicHttpHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}
