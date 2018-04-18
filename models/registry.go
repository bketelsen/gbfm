package models

import (
	"fmt"

	"github.com/jinzhu/inflection"
)

type registryFuncs struct {
	// a function that returns an empty model
	empty func() IDer
	// a function that returns a models that is full of sample data.
	//
	// useful for testing
	sample func() IDer
	// a function that returns a list of empty models
	list func() interface{}
}

// for each type name, provide a function that returns an empty type and an empty list of
// that type
var registry = map[string]*registryFuncs{}

// EmptyFromRegistry returns a new empty model
func EmptyFromRegistry(name string) (IDer, error) {
	funcs, ok := registry[name]
	if !ok {
		singular := inflection.Singular(name)
		funcs, ok = registry[singular]
		if !ok {
			return nil, fmt.Errorf("unknown model %s / %s", name, singular)
		}
	}
	return funcs.empty(), nil
}

// SampleFromRegistry returns a new model that is full of sample data.
//
// useful for testing
func SampleFromRegistry(name string) (IDer, error) {
	funcs, ok := registry[name]
	if !ok {
		singular := inflection.Singular(name)
		funcs, ok = registry[singular]
		if !ok {
			return nil, fmt.Errorf("unknown model %s / %s", name, singular)
		}
	}
	return funcs.sample(), nil
}

// EmptyListFromRegistry returns a new list of models
func EmptyListFromRegistry(name string) (interface{}, error) {
	funcs, ok := registry[name]
	if !ok {
		singular := inflection.Singular(name)
		funcs, ok = registry[singular]
		if !ok {
			return nil, fmt.Errorf("unknown model %s / %s", name, singular)
		}
	}
	return funcs.list(), nil
}
