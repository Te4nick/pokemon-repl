package callbacks

import (
	"encoding/json"
	"fmt"
	"github.com/chrxn1c/pokemon-repl/internal/user_context"
	"io"
	"log"
	"net/http"
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
		"3) type \"map\" command")

	return "", nil
}

func MapCallback(ctx *user_context.UserContext) (string, error) {
	currentURL := fmt.Sprintf("https://pokeapi.co/api/v2/location?limit=20&offset=%d", ctx.APIoffset)
	result, err := http.Get(currentURL)

	if err != nil {
		log.Fatalf("Failed to fetch response when doing $map command, user context: %v\nerr: %v\nURL: %v\n", ctx, err, currentURL)
	}

	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Fatalf("Failed to parse response when have doing $map command, user context: %v\nerr: %v\nURL: %v\n", ctx, err, currentURL)
	}
	err = result.Body.Close()
	if err != nil {
		log.Fatalf("Failed to close body of response when have received $map command, user context: %v\nerr: %v\nURL: %v\n", ctx, err, currentURL)
	}

	if result.StatusCode >= 400 {
		log.Fatalf("Response of $map failed with status code: %d and\nbody: %s\n", result.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	ctx.APIoffset += 20

	type MapResponse struct {
		Count    int    `json:"-"`
		Next     string `json:"-"`
		Previous any    `json:"-"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"-"`
		} `json:"results"`
	}

	mapResponse := MapResponse{}
	err = json.Unmarshal(body, &mapResponse)
	if err != nil {
		return "", err
	}

	toUserResponse := ""
	for _, location := range mapResponse.Results {
		toUserResponse += fmt.Sprintf("%v\n", location.Name)
	}
	return toUserResponse, nil
}
