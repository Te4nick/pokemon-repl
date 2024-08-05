package callbacks

import (
	"encoding/json"
	"fmt"
	"github.com/chrxn1c/pokemon-repl/internal/user_context"
	"log"
	"os"
)

type CallbackInterface interface {
	Call(data string) (output string, err error)
}

func ExitCallback(ctx *user_context.UserContext) (string, error) {
	fmt.Println("\nIt's been a pleasure to have you onboard! Thanks for using this application")
	os.Exit(0)
	return "", nil
}

func HelpCallback(ctx *user_context.UserContext) (string, error) {
	fmt.Println("\nFor now you can do the following:\n" +
		"1) type \"help\" command \n" +
		"2) type \"exit\" command \n" +
		"3) type \"map\" command \n" +
		"4) type \"mapb\" command")

	return "", nil
}

func MapCallback(ctx *user_context.UserContext) (string, error) {
	ctx.APIoffset += 20
	currentURL := fmt.Sprintf("https://pokeapi.co/api/v2/location?limit=20&offset=%d", ctx.APIoffset)

	body, err := makeAPIRequestAndProcessErrors(ctx, currentURL)

	apiResponse := mapResponse{}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return "", err
	}

	toUserResponse := ""
	for _, location := range apiResponse.Results {
		toUserResponse += fmt.Sprintf("%v\n", location.Name)
	}
	return toUserResponse, nil
}

func MapbCallback(ctx *user_context.UserContext) (string, error) {
	ctx.APIoffset -= 20
	currentURL := fmt.Sprintf("https://pokeapi.co/api/v2/location?limit=20&offset=%d", ctx.APIoffset)
	if ctx.APIoffset < 0 {
		log.Fatalf("Cannot traverse previous locations since you are at the very beginning ($mapb command), user context: %v\nURL: %v\n", ctx, currentURL)
	}

	body, err := makeAPIRequestAndProcessErrors(ctx, currentURL)

	apiResponse := mapResponse{}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return "", err
	}

	toUserResponse := ""
	for _, location := range apiResponse.Results {
		toUserResponse += fmt.Sprintf("%v\n", location.Name)
	}
	return toUserResponse, nil
}
