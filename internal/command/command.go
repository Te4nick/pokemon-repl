package command

type Command struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Callback
}
