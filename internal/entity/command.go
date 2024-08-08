package entity

type Callback func(ctx *UserContext) (output string, err error)

type Command struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Callback
}
