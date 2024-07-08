package main

import (
	application "github.com/chrxn1c/pokemon-repl/internal/application"
)

func main() {
	app := application.PokemonApplication{}
	err := app.Run()
	if err != nil {
		return
	}
}
