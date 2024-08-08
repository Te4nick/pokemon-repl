package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/chrxn1c/pokemon-repl/pkg/cache"
)

const baseURL string = "https://pokeapi.co/api/v2"

var c *cache.Cache = cache.NewCache(cache.CacheDefaultTTL)

func Fetch(endpoint string, out interface{}) error {
	var apiURL string = ""
	if strings.Contains(endpoint, baseURL) {
		apiURL = endpoint
		endpoint = apiURL[len(baseURL):]
	}

	data := c.Get(endpoint)
	if data != nil {
		reflect.ValueOf(out).Elem().Set(reflect.ValueOf(data).Elem())
		return nil
	}

	var err error
	if apiURL == "" {
		apiURL = baseURL
		sep := ""
		if endpoint[0] != '/' {
			sep = "/"
		}
		apiURL += sep + endpoint
	}

	resp, err := http.Get(apiURL)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("abnormal response status code")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		return err
	}

	c.Set(endpoint, out)
	return nil
}
