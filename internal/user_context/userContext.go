package user_context

import (
	"github.com/chrxn1c/pokemon-repl/internal/user_context/pokemon"
	"github.com/chrxn1c/pokemon-repl/pkg/cache"
)

type UserContext struct {
	CaughtPokemons []pokemon.Pokemon
	APIoffset      int64
	Cache          *cache.Cache
}
