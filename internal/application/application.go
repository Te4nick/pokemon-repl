package application

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chrxn1c/pokemon-repl/internal/command"
	"github.com/chrxn1c/pokemon-repl/internal/pokecache"
	"github.com/chrxn1c/pokemon-repl/internal/user_context"
	"github.com/chrxn1c/pokemon-repl/internal/user_context/pokemon"
	"github.com/chrxn1c/pokemon-repl/internal/utils"
)

type Application interface {
	Run() error
}

type PokemonApplication struct {
	userContext    *user_context.UserContext
	contentManager *utils.ContentManager
	commander      *command.Commander
}

func (app *PokemonApplication) initializeComponents() error {
	app.userContext = &user_context.UserContext{
		APIoffset:      -20, // $ map will increase offset by 20 first and then inspect the given offset
		CaughtPokemons: []pokemon.Pokemon{},
		Cache:          pokecache.NewCache(5 * time.Second),
	}

	var err error
	app.contentManager, err = utils.NewContentManager("en_EN")
	if err != nil {
		return err
	}

	app.commander = command.NewCommander(app.contentManager.Commands)

	return nil
}

func (app *PokemonApplication) printWelcomeMessage() error {
	fmt.Println("Welcome aboard! If you don't have a faintest idea of what to do, type in \"help\" command")
	return nil
}

func (app *PokemonApplication) takeCommand(scanner *bufio.Scanner) (string, error) {
	fmt.Print("pokedex:$ ")
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		fmt.Println("\nError occurred while scanning user input")
		log.Fatal(err)
		return "", err
	}
	stringCommandRepresentation := scanner.Text()

	return stringCommandRepresentation, nil
}

func (app *PokemonApplication) printResultOfCommand(result string) error {
	fmt.Println(result)
	return nil
}

func (app *PokemonApplication) Run() error {
	err := app.initializeComponents()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(os.Stdin)
	err = app.printWelcomeMessage()
	if err != nil {
		return err
	}

	for {
		inferredCommand, err := app.takeCommand(scanner)
		if err != nil {
			return err
		}

		outputData, err := app.commander.Exec(inferredCommand, app.userContext)
		if err != nil {
			return err
		}

		err = app.printResultOfCommand(outputData)
		if err != nil {
			return err
		}
	}
}
