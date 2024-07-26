package callbacks

import (
	"fmt"
	"github.com/chrxn1c/pokemon-repl/internal/user_context"
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
		"2) type \"exit\" command")

	return "", nil
}
