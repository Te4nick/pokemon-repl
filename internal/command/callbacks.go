package command

import (
	"fmt"
	"os"

	"github.com/chrxn1c/pokemon-repl/internal/user_context"
	"github.com/chrxn1c/pokemon-repl/pkg/api"
)

type Callback func(ctx *user_context.UserContext) (output string, err error)

func ExitCallback(ctx *user_context.UserContext) (string, error) {
	fmt.Println("\nIt's been a pleasure to have you onboard! Thanks for using this application")
	os.Exit(0)
	return "", nil
}

func MapCallback(ctx *user_context.UserContext) (string, error) {
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

func MapbCallback(ctx *user_context.UserContext) (string, error) {
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
