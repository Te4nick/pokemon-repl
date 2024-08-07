package main

import (
	application "github.com/chrxn1c/pokemon-repl/internal/application"
)

func main() {
	app, err := application.New()
	if err != nil {
		panic(err)
	}

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
