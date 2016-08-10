package redirect

import (
	"github.com/edio/m1w/storage"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Init(r *mux.Router) {
	r.Methods(http.MethodPost).PathPrefix("/{key}").HandlerFunc(AddRedirect)
	r.Methods(http.MethodGet).PathPrefix("/{key}").HandlerFunc(GoRedirect)
}

func GoRedirect(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	location, err := storage.ResolveLocation(key)
	if err == nil {
		w.Header().Set("Location", *location)
		w.WriteHeader(http.StatusFound)
	} else if err == storage.ErrKeyNotFound {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("x-m1w-error", err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}

func AddRedirect(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	body, _ := ioutil.ReadAll(r.Body)
	url, err := url.Parse(string(body))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = storage.Add(key, url)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.Header().Set("x-m1w-error", err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}
