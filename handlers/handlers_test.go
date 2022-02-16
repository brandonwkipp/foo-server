package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"go-foo-server/foo"
)

// PopulateTestFoo is a helper function that cleanups existing any existing Foo and creates a new test Foo
func PopulateTestFoo() foo.Foo {
	// Cleanup any existing Foo
	foo.ResetFoos()

	existingFoo := foo.Foo{Id: uuid.New().String(), Name: "Test"}
	foo.Foos[existingFoo.Id] = existingFoo

	return existingFoo
}

// TestCreateFooHandlerInvalidData tests the result of CreateFooHandler when the "name" property in the payload is not specified
func TestCreateFooHandlerInvalidData(t *testing.T) {
	// Bogus json without "name" property
	bodyJson, err := json.Marshal(map[string]string{"notname": "Test"})
	if err != nil {
		t.Errorf("Error marshalling JSON: %s", err)
	}

	// Create request and record the response
	request, _ := http.NewRequest(http.MethodPost, "/foo", bytes.NewReader(bodyJson))
	response := httptest.NewRecorder()
	HandleCreateFoo(response, request)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

// TestCreateFooHandlerMissingData tests the result of CreateFooHandler when the payload is empty
func TestCreateFooHandlerMissingData(t *testing.T) {
	// Create request and record the response
	request, _ := http.NewRequest(http.MethodPost, "/foo", nil)
	response := httptest.NewRecorder()
	HandleCreateFoo(response, request)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

// TestCreateFooHandlerSuccess tests the result of CreateFooHandler when the payload is valid
func TestCreateFooHandlerSuccess(t *testing.T) {
	name := "Test"

	bodyJson, err := json.Marshal(map[string]string{"name": name})
	if err != nil {
		t.Errorf("Error marshalling JSON: %s", err)
	}

	// Create request and record the response
	request, _ := http.NewRequest(http.MethodPost, "/foo", bytes.NewReader(bodyJson))
	response := httptest.NewRecorder()
	HandleCreateFoo(response, request)

	// Unmarshal to run assertions on the response
	var f foo.Foo
	err = json.Unmarshal(response.Body.Bytes(), &f)
	if err != nil {
		t.Errorf("Error unmarshalling JSON: %s", err)
	}

	// Assertions
	assert.Equal(t, response.Code, http.StatusOK)
	assert.Equal(t, f.Name, name)
	assert.NotNil(t, f.Id)
}

// TestDeleteFooHandlerNoMatch tests the result of DeleteFooHandler when a match with the supplied Id is not found
func TestDeleteFooHandlerNoMatch(t *testing.T) {
	vars := map[string]string{
		"id": uuid.New().String(),
	}

	// Create request, set vars, and record the response
	request, _ := http.NewRequest(http.MethodDelete, "/foo/"+vars["id"], nil)
	request = mux.SetURLVars(request, vars)
	response := httptest.NewRecorder()
	HandleDeleteFoo(response, request)

	// Assertions
	assert.Equal(t, http.StatusNotFound, response.Code)
}

// TestDeleteFooHandlerSuccess tests the result of DeleteFooHandler when a match with the supplied Id is found
func TestDeleteFooHandlerSuccess(t *testing.T) {
	existingFoo := PopulateTestFoo()

	vars := map[string]string{
		"id": existingFoo.Id,
	}

	// Create request, set vars, and record the response
	request, _ := http.NewRequest(http.MethodDelete, "/foo/"+existingFoo.Id, nil)
	request = mux.SetURLVars(request, vars)
	response := httptest.NewRecorder()
	HandleDeleteFoo(response, request)

	// Assertions
	assert.Equal(t, http.StatusNoContent, response.Code)
}

// TestGetFooHandlerNoMatch tests the result of GetFooHandler when a match with the supplied Id is not found
func TestGetFooHandlerNoMatch(t *testing.T) {
	vars := map[string]string{
		"id": uuid.New().String(),
	}

	// Create request, set vars, and record the response
	request, _ := http.NewRequest(http.MethodGet, "/foo/"+vars["id"], nil)
	request = mux.SetURLVars(request, vars)
	response := httptest.NewRecorder()
	HandleGetFoo(response, request)

	// Assertions
	assert.Equal(t, http.StatusNotFound, response.Code)
}

// TestGetFooHandlerSuccess tests the result of GetFooHandler when a match with the supplied Id is found
func TestGetFooHandlerSuccess(t *testing.T) {
	existingFoo := PopulateTestFoo()

	vars := map[string]string{
		"id": existingFoo.Id,
	}

	// Create request, set vars, and record the response
	request, _ := http.NewRequest(http.MethodGet, "/foo/"+existingFoo.Id, nil)
	request = mux.SetURLVars(request, vars)
	response := httptest.NewRecorder()
	HandleGetFoo(response, request)

	var f foo.Foo
	err := json.Unmarshal(response.Body.Bytes(), &f)
	if err != nil {
		t.Errorf("Error unmarshalling JSON: %s", err)
	}

	// Assertions
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, f.Id, existingFoo.Id)
	assert.Equal(t, f.Name, existingFoo.Name)
}
