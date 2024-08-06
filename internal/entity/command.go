package entity

import (
	"github.com/chrxn1c/pokemon-repl/internal/user_context"
)

type Callback func(ctx *user_context.UserContext) (output string, err error)

type Command struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Callback
}
