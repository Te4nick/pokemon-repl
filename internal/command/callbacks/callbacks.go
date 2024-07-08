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
	fmt.Println("It's been a pleasure to have you on board! Thanks for using this application")
	os.Exit(0)
	return "", nil
}

func HelpCallback(ctx *user_context.UserContext) (string, error) {
	fmt.Println("For now you can do the following:\n" +
		"1) type \"help\" command " +
		"2) type \"exit\" command")

	return "", nil
}
