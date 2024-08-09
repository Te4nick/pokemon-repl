package application

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chrxn1c/pokemon-repl/internal/command"
	"github.com/chrxn1c/pokemon-repl/internal/utils"
	"github.com/chrxn1c/pokemon-repl/pkg/pokectx"
)

type Application interface {
	Run() error
}

type PokemonApplication struct {
	context        *pokectx.Te4nickPokeCTX
	contentManager *utils.ContentManager
	commander      *command.Commander
}

func New() (*PokemonApplication, error) {
	pokeCTX := pokectx.New()
	pokectx.SetNum(pokeCTX, -20, "api", "location", "offset")

	var err error
	contentManager, err := utils.NewContentManager("en_EN")
	if err != nil {
		return nil, err
	}

	commander := command.NewCommander(contentManager.Commands)

	return &PokemonApplication{
		context:        pokeCTX,
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

		outputData, err := app.commander.Exec(inferredCommand, argument, app.context)
		if err != nil {
			return err
		}

		err = app.printResultOfCommand(outputData)
		if err != nil {
			return err
		}
	}
}
