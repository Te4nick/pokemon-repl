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
		return "Detected argument to exit command which is not supported", nil
	}
	fmt.Println("\nIt's been a pleasure to have you onboard! Thanks for using this application")
	os.Exit(0)
	return "", nil
}

func MapCallback(ctx *pokectx.Te4nickPokeCTX, arg string) (string, error) {
	if len(arg) != 0 {
		return "Detected argument to map command which is not supported", nil
	}

	apiOffset := pokectx.GetOrDefaultNum(ctx, 0, "api", "location", "offset")
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
	if len(arg) != 0 {
		return "Detected argument to map command which is not supported", nil
	}

	apiOffset := pokectx.GetOrDefaultNum(ctx, 0, "api", "location", "offset")
	apiOffset -= 20
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

func ExploreCallback(ctx *pokectx.Te4nickPokeCTX, arg string) (string, error) {
	if len(arg) == 0 || strings.Contains(arg, " ") {
		return "You need to pass area-name as an argument to explore it.", nil
	}

	locationArea, err := api.LocationArea(arg)
	if err != nil {
		if statusErr, ok := err.(api.HTTPStatusError); ok && statusErr.StatusCode == 404 {
			return "Location Area not found: " + arg, nil
		}
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
	if len(strings.Fields(arg)) != 1 {
		return "This command requires an argument, virtually one.", nil
	}

	baseChance := 0.4 + rand.Float32()/2.0

	pokemon, err := api.Pokemon(arg)
	if err != nil {
		if statusErr, ok := err.(api.HTTPStatusError); ok && statusErr.StatusCode == 404 {
			return "Pokemon not found: " + arg, nil
		}
		return "", err
	}

	if baseChance-float32(pokemon.BaseExperience)/1000.0 < 0.5 {
		return "Couldn't catch " + pokemon.Name + "... Try again!", nil
	}

	ctx.SetKey("pokedex", pokemon.Name)
	return "Congratulations! You've caught " + pokemon.Name + "!", nil
}

func InspectCallback(ctx *pokectx.Te4nickPokeCTX, arg string) (string, error) {
	if len(strings.Fields(arg)) != 1 {
		return "This command requires an argument, virtually one.", nil
	}

	_, found := ctx.Get("pokedex", arg)
	if !found {
		return "You have not caught that pokemon", nil
	}

	pokemonInfo, err := api.Pokemon(arg)
	if err != nil {
		if statusErr, ok := err.(api.HTTPStatusError); ok && statusErr.StatusCode == 404 {
			return "Pokemon not found: " + arg, nil
		}
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

	caughtPokemons, _ := ctx.GetKeys("pokedex")

	msg := "Your Pokedex:\n"
	for _, pokemonName := range caughtPokemons {
		msg += "  - " + pokemonName + "\n"
	}

	return msg, nil
}
