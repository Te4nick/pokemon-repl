package api

import (
	"encoding/json"
	"io"
	"net/http"
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
		err := json.Unmarshal(data, out)
		if err != nil {
			return err
		}
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
		return HTTPStatusError{StatusCode: resp.StatusCode, URL: apiURL}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		return err
	}

	c.Set(endpoint, body)
	return nil
}
