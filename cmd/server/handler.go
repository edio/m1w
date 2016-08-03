package main

import (
	"github.com/edio/monolith"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

func AdminUi(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	w.Write([]byte("Under construction:" + key))
	w.WriteHeader(http.StatusFound)
}

func GoRedirect(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	location, err := monolith.ResolveLocation(key)
	if err == nil {
		w.Header().Set("Location", *location)
		w.WriteHeader(http.StatusFound)
	} else if err == monolith.ErrKeyNotFound {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("x-go-error", err.Error())
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

	err = monolith.Add(key, url)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.Header().Set("x-go-error", err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}
