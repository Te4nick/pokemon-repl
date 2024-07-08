package application

import (
	"bufio"
	"errors"
	"github.com/chrxn1c/pokemon-repl/internal/command"
	"github.com/chrxn1c/pokemon-repl/internal/user_context"
	"github.com/chrxn1c/pokemon-repl/internal/user_context/pokemon"
)

type Application interface {
	initializeComponents()
	printWelcomeMessage(writer *bufio.Writer) error
	takeCommand(scanner *bufio.Scanner) (command.Command, error)
	evaluateCommand(cmd *command.Command) (string, error)
	printResultOfCommand(result string, writer *bufio.Writer) error
	Run() error
}

type PokemonApplication struct {
	userContext user_context.UserContext
}

func (app *PokemonApplication) initializeComponents() {
	app.userContext = user_context.UserContext{
		APIoffset:      0,
		CaughtPokemons: []pokemon.Pokemon{},
	}
}

func (app *PokemonApplication) printWelcomeMessage(writer *bufio.Writer) error {
	_, err := writer.WriteString("Welcome aboard! If you don't have a faintest idea of what to do, type in \"help\" command")
	if err != nil {
		return err
	}
	return nil
}

func (app *PokemonApplication) takeCommand(scanner *bufio.Scanner) (command.Command, error) {
	stringCommandRepresentation := scanner.Text()

	cmd, ok := command.Commands[stringCommandRepresentation]

	if !ok {
		return command.Command{}, errors.New("such command is not supported")
	}

	return cmd, nil
}

func (app *PokemonApplication) evaluateCommand(cmd command.Command) (string, error) {
	outputData, err := cmd.Callback(&app.userContext)
	if err != nil {
		return outputData, err
	}
	return outputData, nil
}

func (app *PokemonApplication) printResultOfCommand(result string, writer *bufio.Writer) error {
	_, err := writer.WriteString(result)
	if err != nil {
		return err
	}
	return nil
}

func (app *PokemonApplication) Run() error {
	app.initializeComponents()
	scanner := bufio.Scanner{}
	writer := bufio.Writer{}

	err := app.printWelcomeMessage(&writer)
	if err != nil {
		return err
	}

	for {
		inferredCommand, err := app.takeCommand(&scanner)
		if err != nil {
			return err
		}

		outputData, err := app.evaluateCommand(inferredCommand)

		err = app.printResultOfCommand(outputData, &writer)
		if err != nil {
			return err
		}
	}
}
