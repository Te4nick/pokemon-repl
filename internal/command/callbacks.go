package command

import (
	"fmt"
	"os"

	"github.com/chrxn1c/pokemon-repl/internal/entity"
	"github.com/chrxn1c/pokemon-repl/pkg/api"
)

func ExitCallback(ctx *entity.UserContext, arg string) (string, error) {
	if len(arg) > 0 {
		fmt.Println("Detected argument to help command which is not supported")
		return "", nil
	}
	fmt.Println("\nIt's been a pleasure to have you onboard! Thanks for using this application")
	os.Exit(0)
	return "", nil
}

func MapCallback(ctx *entity.UserContext, arg string) (string, error) {
	if len(arg) > 0 {
		fmt.Println("Detected argument to map command which is not supported")
		return "", nil
	}

	ctx.APIoffset += 20
	endpoint := fmt.Sprintf("location?limit=20&offset=%d", ctx.APIoffset)

	locations, err := api.Resource(endpoint)
	if err != nil {
		return "", err
	}

	toUserResponse := ""
	for _, location := range locations.Results {
		toUserResponse += fmt.Sprintf("%v\n", location.Name)
	}
	return toUserResponse, nil
}

func MapbCallback(ctx *entity.UserContext, arg string) (string, error) {
	if len(arg) > 0 {
		fmt.Println("Detected argument to map command which is not supported")
		return "", nil
	}

	ctx.APIoffset -= 20
	endpoint := fmt.Sprintf("location?limit=20&offset=%d", ctx.APIoffset)

	locations, err := api.Resource(endpoint)
	if err != nil {
		return "", err
	}

	toUserResponse := ""
	for _, location := range locations.Results {
		toUserResponse += fmt.Sprintf("%v\n", location.Name)
	}
	return toUserResponse, nil
}

func ExploreCallback(ctx *entity.UserContext, arg string) (string, error) {
	if len(arg) == 0 {
		fmt.Println("You need to pass area-name as an argument to explore it.")
		return "", nil
	}

	locationArea, err := api.LocationArea(arg)
	if err != nil {
		return "", err
	}

	toUserResponse := fmt.Sprintf("Exploring %v ...\n", arg)
	foundPokemonsAnnouncement := "Found Pokemon:\n"
	foundPokemons := ""

	for _, pokemonEncounter := range locationArea.PokemonEncounters {
		foundPokemons += fmt.Sprintf("  - %v\n", pokemonEncounter.Pokemon.Name)
	}

	if len(foundPokemons) == 0 {
		foundPokemonsAnnouncement = "Haven't found any pokemons within given area.\n"
	}

	toUserResponse += foundPokemonsAnnouncement + foundPokemons

	return toUserResponse, nil
}
