package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rhecl/project-x/pkg/mongo"
)

func healthz(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "OK\n")
}

func main() {
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/healthz", healthz).Methods("GET")

	_, err := mongo.Connect()

	if err != nil {
		panic(err.Error())
	}

	http.ListenAndServe(":3001", handlers.CORS(headers, methods, origins)(router))
}
