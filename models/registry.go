package models

import (
	"fmt"

	"github.com/jinzhu/inflection"
)

// for each type name, provide a function that returns an empty type and an empty list of
// that type
var registry = map[string]func() (IDer, interface{}){}

// EmptyFromRegistry returns a new model
func EmptyFromRegistry(name string) (IDer, error) {
	fn, ok := registry[name]
	if !ok {
		singular := inflection.Singular(name)
		fn, ok = registry[singular]
		if !ok {
			return nil, fmt.Errorf("unknown model %s / %s", name, singular)
		}
	}
	single, _ := fn()
	return single, nil
}

// EmptyListFromRegistry returns a new list of models
func EmptyListFromRegistry(name string) (interface{}, error) {
	fn, ok := registry[name]
	if !ok {
		singular := inflection.Singular(name)
		fn, ok = registry[singular]
		if !ok {
			return nil, fmt.Errorf("unknown model %s / %s", name, singular)
		}
	}
	_, list := fn()
	return list, nil
}
