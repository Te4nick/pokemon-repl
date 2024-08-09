package entity

import "github.com/chrxn1c/pokemon-repl/pkg/pokectx"

type Callback func(ctx *pokectx.PokeCTX, arg string) (string, error)

type Command struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Callback
}
