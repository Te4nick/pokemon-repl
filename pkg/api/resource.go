package api

import "github.com/chrxn1c/pokemon-repl/pkg/entity"

func Resource(endpoint string) (entity.Resource, error) {
	var out entity.Resource
	err := Fetch(endpoint, &out)
	return out, err
}
