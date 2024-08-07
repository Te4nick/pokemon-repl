package application

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/chrxn1c/pokemon-repl/internal/command"
	"github.com/chrxn1c/pokemon-repl/internal/pokecache"
	"github.com/chrxn1c/pokemon-repl/internal/user_context"
	"github.com/chrxn1c/pokemon-repl/internal/user_context/pokemon"
	"log"
	"os"
	"strings"
	"time"
)

type Application interface {
	Run() error
}

type PokemonApplication struct {
	userContext user_context.UserContext
}

func (app *PokemonApplication) initializeComponents() {
	app.userContext = user_context.UserContext{
		APIoffset:      -20, // $ map will increase offset by 20 first and then inspect the given offset
		CaughtPokemons: []pokemon.Pokemon{},
		Cache:          pokecache.NewCache(5 * time.Second),
	}
}

func (app *PokemonApplication) printWelcomeMessage() error {
	fmt.Println("Welcome aboard! If you don't have a faintest idea of what to do, type in \"help\" command")
	return nil
}

func (app *PokemonApplication) takeCommand(scanner *bufio.Scanner) (inferredCommand command.Command, argument string, err error) {
	fmt.Print("pokedex:$ ")
	scanner.Scan()
	err = scanner.Err()
	if err != nil {
		fmt.Println("\nError occurred while scanning user input")
		log.Fatal(err)
	}
	stringCommandAndArgumentRepresentation := scanner.Text()
	splitString := strings.Split(stringCommandAndArgumentRepresentation, " ")

	inferredCommand, ok := command.Commands[splitString[0]]

	if !ok {
		return command.Command{}, "", errors.New("such command is not supported")
	}

	if len(splitString) > 1 {
		argument = splitString[1]
	}

	return inferredCommand, argument, nil
}

func (app *PokemonApplication) evaluateCommand(cmd command.Command, arg string) (string, error) {
	outputData, err := cmd.Callback(&app.userContext, arg)
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
		inferredCommand, argument, err := app.takeCommand(scanner)
		if err != nil {
			app.unknownCommand()
			continue
		}

		outputData, err := app.evaluateCommand(inferredCommand, argument)

		err = app.printResultOfCommand(outputData)
		if err != nil {
			return err
		}
	}
}
