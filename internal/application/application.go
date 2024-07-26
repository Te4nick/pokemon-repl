package application

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/chrxn1c/pokemon-repl/internal/command"
	"github.com/chrxn1c/pokemon-repl/internal/user_context"
	"github.com/chrxn1c/pokemon-repl/internal/user_context/pokemon"
	"log"
	"os"
)

type Application interface {
	initializeComponents()
	printWelcomeMessage() error
	takeCommand(scanner *bufio.Scanner) (command.Command, error)
	evaluateCommand(cmd *command.Command) (string, error)
	printResultOfCommand(result string) error
	unknownCommand()
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

func (app *PokemonApplication) printWelcomeMessage() error {
	fmt.Println("Welcome aboard! If you don't have a faintest idea of what to do, type in \"help\" command")
	return nil
}

func (app *PokemonApplication) takeCommand(scanner *bufio.Scanner) (command.Command, error) {
	fmt.Print("pokedex:$ ")
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		fmt.Println("\nError occurred while scanning user input")
		log.Fatal(err)
	}
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

func (app *PokemonApplication) printResultOfCommand(result string) error {
	fmt.Println(result)
	return nil
}

func (app *PokemonApplication) unknownCommand() {
	fmt.Println("Given command is not supported. Use \"help\" if necessary.")
}

func (app *PokemonApplication) Run() error {
	app.initializeComponents()
	scanner := bufio.NewScanner(os.Stdin)

	err := app.printWelcomeMessage()
	if err != nil {
		return err
	}

	for {
		inferredCommand, err := app.takeCommand(scanner)
		if err != nil {
			app.unknownCommand()
			continue
		}

		outputData, err := app.evaluateCommand(inferredCommand)

		err = app.printResultOfCommand(outputData)
		if err != nil {
			return err
		}
	}
}
