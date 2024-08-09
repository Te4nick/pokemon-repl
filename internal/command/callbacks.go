package command

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/chrxn1c/pokemon-repl/pkg/api"
	"github.com/chrxn1c/pokemon-repl/pkg/pokectx"
)

func ExitCallback(ctx *pokectx.Te4nickPokeCTX, arg string) (string, error) {
	if len(arg) > 0 {
		fmt.Println("Detected argument to help command which is not supported")
		return "", nil
	}
	fmt.Println("\nIt's been a pleasure to have you onboard! Thanks for using this application")
	os.Exit(0)
	return "", nil
}

func MapCallback(ctx *pokectx.Te4nickPokeCTX, arg string) (string, error) {
	if len(arg) > 0 {
		fmt.Println("Detected argument to map command which is not supported")
		return "", nil
	}

	apiOffset := pokectx.GetDefaultNum(ctx, 0, "api", "location", "offset")
	apiOffset += 20
	pokectx.SetNum(ctx, apiOffset, "api", "location", "offset")
	endpoint := fmt.Sprintf("location-area?limit=20&offset=%d", apiOffset)

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

func MapbCallback(ctx *pokectx.Te4nickPokeCTX, arg string) (string, error) {
	if len(arg) == 0 {
		fmt.Println("Detected argument to map command which is not supported")
		return "", nil
	}

	apiOffset := pokectx.GetDefaultNum(ctx, 0, "api", "location", "offset")
	apiOffset -= 20
	pokectx.SetNum(ctx, apiOffset, "api", "location", "offset")
	endpoint := fmt.Sprintf("location?limit=20&offset=%d", apiOffset)

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

func ExploreCallback(ctx *pokectx.Te4nickPokeCTX, arg string) (string, error) {
	if len(arg) == 0 || strings.Contains(arg, " ") {
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

func CatchCallback(ctx *pokectx.Te4nickPokeCTX, arg string) (string, error) {
	if len(arg) == 0 || strings.Contains(arg, " ") {
		fmt.Println("You need to pass pokemon name as an argument to catch it.")
		return "", nil
	}

	baseChance := 0.3 + rand.Float32()/2.0

	pokemon, err := api.Pokemon(arg)
	if err != nil {
		return "", err
	}

	if baseChance-float32(pokemon.BaseExperience)/1000.0 < 0.5 {
		return "Couldn't catch " + pokemon.Name + "... Try again!", nil
	}

	ctx.Set("pokedex", pokemon.Name)
	return "Congratulations! You've caught " + pokemon.Name + "!", nil
}

func InspectCallback(ctx *pokectx.Te4nickPokeCTX, arg string) (string, error) {
	if len(arg) == 0 || strings.Contains(arg, " ") {
		fmt.Println("This command requires an argument, virtually one.")
		return "", nil
	}

	caughtPokemons, _ := ctx.Get("pokedex")
	if !strings.Contains(caughtPokemons, arg) {
		return "You have not caught that pokemon\n", nil
	}

	// TODO: inspect pokemon via te4nick structure
	pokemonInfo, err := api.Pokemon(arg)
	if err != nil {
		return "Error when making request to the API", err
	}

	toUserResponse := ""
	toUserResponse += "Name: " + pokemonInfo.Name + "\n"
	toUserResponse += "Height: " + strconv.Itoa(pokemonInfo.Height) + "\n"
	toUserResponse += "Weight: " + strconv.Itoa(pokemonInfo.Weight) + "\n"
	toUserResponse += "Stats: " + "\n"

	for _, characteristic := range pokemonInfo.Stats {
		toUserResponse += "  - " + characteristic.Stat.Name + ": " + strconv.Itoa(characteristic.BaseStat) + "\n"
	}

	toUserResponse += "Types: " + "\n"

	for _, pokemonType := range pokemonInfo.Types {
		toUserResponse += "  - " + pokemonType.Type.Name + " \n"
	}

	return toUserResponse, nil
}

func PokedexCallback(ctx *pokectx.Te4nickPokeCTX, arg string) (string, error) {
	if len(arg) != 0 {
		fmt.Println("This command does not require any arguments.")
		return "", nil
	}

	caughtPokemons, _ := ctx.Get("pokedex")

	return caughtPokemons, nil
}
