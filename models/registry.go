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
	// a function that returns a _pointer_ to an empty list of a model. it really needs
	// to return a pointer, or you'll get crazy errors when you use it to get
	// a list of models from the DB.
	//
	// for example, the registry entry for an Episode must be this:
	//
	//	func() interface{} { return new([]Episode) }
	list func() Lister
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

// EmptyListFromRegistry returns a new empty list of a model with name.
//
// For example, if you want to call this function to get all Episodes, use the
// following:
//
//	list, err := EmptyListFromRegistry("episodes")
//	if err != nil {
//		panic(err)
//	}
//	err := DB.All(list) // do not use &list here or you'll get SQL errors!
//	if err != nil {
//		panic(err)
//	}
func EmptyListFromRegistry(name string) (Lister, error) {
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
