package callbacks

import (
	"github.com/chrxn1c/pokemon-repl/internal/user_context"
	"os"
)

type CallbackInterface interface {
	Call(data string) (output string, err error)
}

//type ExitCallback struct {
//}

type HelpCallback struct {
}

//func (c *ExitCallback) Call() error {
//	os.Exit(0)
//	return nil
//}

func ExitCallback(ctx *user_context.UserContext) (string, error) {
	os.Exit(0)
	return "", nil
}
