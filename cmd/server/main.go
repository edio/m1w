package main

import (
	"flag"
	"github.com/edio/monolith"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	port := flag.Int("port", 80, "Port at which the http server binds to")
	db := flag.String("db", "monolith.db", "Directory where Monolith database is stored")
	flag.Parse()

	r := mux.NewRouter()
	r.PathPrefix("/_static/").Handler(http.StripPrefix("/_static/", http.FileServer(http.Dir("static"))))
	r.Methods(http.MethodGet).Path("/").HandlerFunc(AdminUi)
	r.Methods(http.MethodGet).Path("/_ui").HandlerFunc(AdminUi)
	r.Methods(http.MethodGet).PathPrefix("/_ui/{key}").HandlerFunc(AdminUi)
	r.Methods(http.MethodPost).PathPrefix("/{key}").HandlerFunc(AddRedirect)
	r.Methods(http.MethodGet).PathPrefix("/{key}").HandlerFunc(GoRedirect)

	err := monolith.Init(*db)
	if err != nil {
		log.Fatal(err)
	}
	defer monolith.Close()

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + strconv.Itoa(*port),
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
