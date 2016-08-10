package ui

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Init(r *mux.Router) {
	r.PathPrefix("/_static/").Handler(http.StripPrefix("/_static/", http.FileServer(http.Dir("static"))))
	r.Methods(http.MethodGet).Path("/").HandlerFunc(AdminUi)
	r.Methods(http.MethodGet).Path("/_ui").HandlerFunc(AdminUi)
	r.Methods(http.MethodGet).PathPrefix("/_ui/{key}").HandlerFunc(AdminUi)
}

func AdminUi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Under construction"))
	w.WriteHeader(http.StatusFound)
}
