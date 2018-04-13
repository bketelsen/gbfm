package models

import (
	"fmt"
)

// for each type name, provide a function that returns an empty type and an empty list of
// that type
var registry = map[string]func() (interface{}, interface{}){}

// EmptyFromRegistry returns a new model
func EmptyFromRegistry(name string) (interface{}, error) {
	fn, ok := registry[name]
	if !ok {
		return nil, fmt.Errorf("unknown model %s", name)
	}
	single, _ := fn()
	return single, nil
}

// EmptyListFromRegistry returns a new list of models
func EmptyListFromRegistry(name string) (interface{}, error) {
	fn, ok := registry[name]
	if !ok {
		return nil, fmt.Errorf("unknown model %s", name)
	}
	_, list := fn()
	return list, nil
}
