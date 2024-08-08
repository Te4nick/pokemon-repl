package entity

import (
	"github.com/chrxn1c/pokemon-repl/pkg/cache"
	"github.com/chrxn1c/pokemon-repl/pkg/entity"
)

type UserContext struct {
	CaughtPokemons []entity.Pokemon
	APIoffset      int64
	Cache          *cache.Cache
}
