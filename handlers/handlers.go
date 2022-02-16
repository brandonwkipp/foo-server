package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"

	"go-foo-server/foo"
)

type Payload struct {
	Name string `json:"name" validate:"required"`
}

// HandleCreateFoo creates a new Foo from the given Payload
func HandleCreateFoo(w http.ResponseWriter, r *http.Request) {
	var p Payload

	if r.Body == nil {
		http.Error(w, "No data provided. Please provide a JSON payload, e.g. {\"name\": \"foo\"}", http.StatusBadRequest)
		return
	}

	// Decode the request body into the Payload struct
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		if err.Error() == "EOF" {
			http.Error(w, "No data provided. Please provide a JSON payload, e.g. {\"name\": \"foo\"}", http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the Payload to ensure the name property is not empty
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		http.Error(w, "Please provide a valid JSON payload, e.g. {\"name\": \"foo\"}", http.StatusBadRequest)
		return
	}

	// Create a new Foo with the given name
	f, err := foo.CreateFoo(p.Name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Return the newly created Foo
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(f)
}

// HandleDeleteFoo deletes a Foo identified by the Id specified in the URL path
func HandleDeleteFoo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Search for the Foo with the given ID and cut it
	fooFoundAndRemoved := foo.DeleteFoo(id)
	if fooFoundAndRemoved {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

// HandleGetFoo returns a Foo identified by the Id specified in the URL path
func HandleGetFoo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	f, err := foo.GetFoo(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(f)
}
