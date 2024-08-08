package api

import "github.com/chrxn1c/pokemon-repl/pkg/entity"

func LocationArea(id string) (result entity.LocationArea, err error) {
	err = Fetch("/location-area/"+id+"/", &result)
	return result, err
}
