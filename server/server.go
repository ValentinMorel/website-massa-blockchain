package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func DefineServer(files map[string][]byte) *http.Server {
	// Declare server and routes
	// It's ok for a prototype
	// Would be better to parse the URL to select the file to display / load
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")

		w.Write(files["index.html"])

	})
	router.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "style")

		w.Write(files["style.css"])

	})
	router.HandleFunc("/script.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		w.Write(files["script.js"])

	})

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv
}
