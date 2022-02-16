package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-foo-server/handlers"
)

func main() {
	port := ":8080"
	router := mux.NewRouter()

	router.HandleFunc("/foo", handlers.HandleCreateFoo).Methods("POST")
	router.HandleFunc("/foo/{id}", handlers.HandleDeleteFoo).Methods("DELETE")
	router.HandleFunc("/foo/{id}", handlers.HandleGetFoo).Methods("GET")

	http.Handle("/", router)

	// Start server
	log.Printf("Server listening on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
