package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var dir string

	flag.StringVar(&dir, "dir", "./build", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	r := mux.NewRouter()
	r.Use(SetHeader)
	r.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	// This will serve files under http://localhost:8000/static/<filename>
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:4567",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Running components server on port 4567")

	log.Fatal(srv.ListenAndServe())
}

func SetHeader(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "Golang/1.18")
		w.Header().Set("Content-Type", "application/javascript")
        next.ServeHTTP(w, r)
    })
}