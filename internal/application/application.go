package application

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chrxn1c/pokemon-repl/internal/command"
	"github.com/chrxn1c/pokemon-repl/internal/entity"
	"github.com/chrxn1c/pokemon-repl/internal/utils"
	"github.com/chrxn1c/pokemon-repl/pkg/cache"
)

type Application interface {
	Run() error
}

type PokemonApplication struct {
	userContext    *entity.UserContext
	contentManager *utils.ContentManager
	commander      *command.Commander
}

func New() (*PokemonApplication, error) {
	userContext := &entity.UserContext{
		APIoffset: -20, // $ map will increase offset by 20 first and then inspect the given offset
		// CaughtPokemons: []entity.Pokemon{}, // TODO: move to context
		Cache: cache.NewCache(5 * time.Second),
	}

	var err error
	contentManager, err := utils.NewContentManager("en_EN")
	if err != nil {
		return nil, err
	}

	commander := command.NewCommander(contentManager.Commands)

	return &PokemonApplication{
		userContext:    userContext,
		contentManager: contentManager,
		commander:      commander,
	}, nil
}

func (app *PokemonApplication) printWelcomeMessage() error {
	fmt.Println("Welcome aboard! If you don't have a faintest idea of what to do, type in \"help\" command")
	return nil
}

func (app *PokemonApplication) takeCommand(scanner *bufio.Scanner) (inferredCommand string, argument string, err error) {

	fmt.Print("pokedex:$ ")
	scanner.Scan()
	err = scanner.Err()
	if err != nil {
		fmt.Println("\nError occurred while scanning user input")
		log.Fatal(err)
		return "", "", err
	}
	stringCommandAndArgumentRepresentation := scanner.Text()
	splitString := strings.Split(stringCommandAndArgumentRepresentation, " ")

	inferredCommand = splitString[0]

	if len(splitString) > 1 {
		argument = splitString[1]
	}

	return inferredCommand, argument, nil
}

func (app *PokemonApplication) printResultOfCommand(result string) error {
	fmt.Println(result)
	return nil
}

func (app *PokemonApplication) Run() error {
	scanner := bufio.NewScanner(os.Stdin)
	err := app.printWelcomeMessage()
	if err != nil {
		return err
	}

	for {
		inferredCommand, argument, err := app.takeCommand(scanner)
		if err != nil {
			return err
		}

		outputData, err := app.commander.Exec(inferredCommand, argument, app.userContext)
		if err != nil {
			return err
		}

		err = app.printResultOfCommand(outputData)
		if err != nil {
			return err
		}
	}
}
