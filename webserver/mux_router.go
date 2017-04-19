package webserver

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/kihoonkim/gogo/user"
	"time"
	"log"
)

var router *mux.Router

func InitRouter() {
	r := mux.NewRouter()
	user.HandlerMapping(r)

	http.Handle("/", r)
	router = r
}

func Run() {
	server := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
