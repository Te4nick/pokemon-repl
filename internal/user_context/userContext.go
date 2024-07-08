package user_context

import (
	"github.com/chrxn1c/pokemon-repl/internal/user_context/pokemon"
)

type UserContext struct {
	CaughtPokemons []pokemon.Pokemon
	APIoffset      int64
}
