package foo

import (
	"errors"

	"github.com/google/uuid"
)

type Foo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Foos is a map of Foos
var Foos map[string]Foo = make(map[string]Foo)

// CreateFoo creates a new Foo for a given name and adds it to the list of Foos
func CreateFoo(name string) (Foo, error) {
	if name != "" {
		f := Foo{
			Id:   uuid.New().String(),
			Name: name,
		}
		// Append to the list of Foos
		Foos[f.Id] = f
		return f, nil
	}
	return Foo{}, errors.New("name property must not be empty")
}

// DeleteFoo cuts a Foo from the Foos map
func DeleteFoo(id string) bool {
	if _, found := Foos[id]; found {
		delete(Foos, id)
		return true
	}
	return false
}

// GetFoo gets a Foo from the list of Foos
func GetFoo(id string) (Foo, error) {
	if _, found := Foos[id]; found {
		return Foos[id], nil
	}
	return Foo{}, errors.New("foo not found")
}

// ResetFoos is a helper function to reset the Foos map
func ResetFoos() {
	Foos = make(map[string]Foo)
}
