package japp

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// InitRouter router
func InitRouter() *mux.Router {
	router := mux.NewRouter()
	dir, _ := os.Getwd()

	html := dir + "/manta/public/"

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, html+"html/index.html")
	})
	router.HandleFunc("/execute", Chain(ExecuteHandler, Method("POST"), Logging()))
	router.HandleFunc("/ws", Chain(WebsocketHandler, Logging()))
	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir(html))))
	return router
}

// Start new server
func Start() *http.Server {

	router := InitRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv
}
