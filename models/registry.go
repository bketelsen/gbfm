package models

import (
	"fmt"
)

var registry = map[string]func() (CoreTemplater, []CoreTemplater){}

// EmptyFromRegistry returns a new model
func EmptyFromRegistry(name string) (CoreTemplater, error) {
	fn, ok := registry[name]
	if !ok {
		return nil, fmt.Errorf("unknown model %s", name)
	}
	single, _ := fn()
	return single, nil
}

// EmptyListFromRegistry returns a new list of models
func EmptyListFromRegistry(name string) ([]CoreTemplater, error) {
	fn, ok := registry[name]
	if !ok {
		return nil, fmt.Errorf("unknown model %s", name)
	}
	_, list := fn()
	return list, nil
}
