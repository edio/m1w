package main

import (
	"flag"
	"github.com/edio/m1w/ui"
	"github.com/edio/m1w/redirect"
	"github.com/edio/m1w/storage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	port := flag.Int("port", 80, "Port to bind to")
	db := flag.String("db", "m1w.db", "Path to database directory")
	enableUi := flag.Bool("ui", false, "Enable UI")
	flag.Parse()

	r := mux.NewRouter()

	redirect.Init(r)
	if *enableUi {
		ui.Init(r)
	}

	err := storage.Init(*db)
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + strconv.Itoa(*port),
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
