package utils

import (
	"encoding/json"
	"os"
	"path"

	"github.com/chrxn1c/pokemon-repl/internal/entity"
)

type ContentManager struct {
	Commands []*entity.Command
}

func NewContentManager(locale string) (*ContentManager, error) {
	localeFilePath := path.Join("assets/locale/", locale+".json")
	content, err := os.Open(localeFilePath)
	if err != nil {
		localeFilePath = "assets/locale/en_EN.json"
		content, err = os.Open(localeFilePath)
	}

	if err != nil {
		return nil, err
	}

	var cm *ContentManager
	err = json.NewDecoder(content).Decode(&cm)
	if err != nil {
		return nil, err
	}

	return cm, nil
}
