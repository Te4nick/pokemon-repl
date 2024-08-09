package api

import "github.com/chrxn1c/pokemon-repl/pkg/entity"

func Pokemon(id string) (result entity.Pokemon, err error) {
	err = Fetch("pokemon/"+id+"/", &result)
	return result, err
}
