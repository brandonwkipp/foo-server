package foo

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// EnsureNoFoosExist is a helper function to assert that no Foos exist.
func EnsureNoFoosExist(t *testing.T) {
	assert.Equal(t, len(Foos), 0)
}

// TestCreateFooInvalidName tests the result of CreateFoo when an empty string is supplied.
func TestCreateFooInvalidName(t *testing.T) {
	EnsureNoFoosExist(t)

	// Attempt to create Foo with empty name
	f, err := CreateFoo("")
	assert.Equal(t, f, Foo{})
	assert.NotNil(t, err, "name property must not be empty")

	ResetFoos()
}

// TestCreateFooSuccess tests the result of CreateFoo when a valid string is supplied.
func TestCreateFooSuccess(t *testing.T) {
	EnsureNoFoosExist(t)

	// Create Foo with proper name
	name := "Test"
	f, err := CreateFoo(name)

	assert.Nil(t, err)
	assert.Equal(t, f.Name, name)
	assert.NotNil(t, f.Id)
	assert.Equal(t, len(Foos), 1)
	assert.Equal(t, Foos[f.Id], f)

	ResetFoos()
}

// TestDeleteFooInvalidId tests the result of DeleteFoo when an invalid Id is supplied.
func TestDeleteFooInvalidId(t *testing.T) {
	EnsureNoFoosExist(t)

	fooIsDeleted := DeleteFoo("")
	assert.False(t, fooIsDeleted)
}

// TestDeleteFooNoMatch tests the result of DeleteFoo when a valid Id is supplied but a match is not found.
func TestDeleteFooNoMatch(t *testing.T) {
	EnsureNoFoosExist(t)

	fooIsDeleted := DeleteFoo("Test")
	assert.False(t, fooIsDeleted)
}

// TestDeleteFooSuccess tests the result of DeleteFoo when a valid Id is supplied and a match is found.
func TestDeleteFooSuccess(t *testing.T) {
	EnsureNoFoosExist(t)

	// Setup up existing Foo
	f := Foo{Id: uuid.New().String()}
	Foos[f.Id] = f

	fooIsDeleted := DeleteFoo(f.Id)
	assert.True(t, fooIsDeleted)
	assert.Equal(t, len(Foos), 0)

	ResetFoos()
}

// TestGetFooInvalidId tests the result of GetFoo when a valid Id is supplied but a match is not found.
func TestGetFooInvalidId(t *testing.T) {
	EnsureNoFoosExist(t)

	// Check empty id returns error
	f, err := GetFoo("")
	assert.Equal(t, f, Foo{})
	assert.EqualError(t, err, "foo not found")
}

// TestGetFooNoMatch tests the result of GetFoo when a valid Id is supplied but a match is not found.
func TestGetFooNoMatch(t *testing.T) {
	EnsureNoFoosExist(t)

	f, err := GetFoo(uuid.New().String())
	assert.Equal(t, f, Foo{})
	assert.EqualError(t, err, "foo not found")
}

// TestGetFooSuccess tests the result of GetFoo when a valid Id is supplied and a match is found.
func TestGetFooSuccess(t *testing.T) {
	EnsureNoFoosExist(t)

	// Setup up existing Foo
	existingFoo := Foo{Id: uuid.New().String(), Name: "Test"}
	Foos[existingFoo.Id] = existingFoo

	f, err := GetFoo(existingFoo.Id)
	assert.Equal(t, f.Id, existingFoo.Id)
	assert.Equal(t, f.Name, existingFoo.Name)
	assert.Nil(t, err)

	ResetFoos()
}
