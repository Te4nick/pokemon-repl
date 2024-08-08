package entity

type Callback func(ctx *UserContext, arg string) (string, error)

type Command struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Callback
}
