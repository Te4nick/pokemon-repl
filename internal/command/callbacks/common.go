package callbacks

import (
	"github.com/chrxn1c/pokemon-repl/internal/user_context"
	"io"
	"log"
	"net/http"
)

func makeAPIRequestAndProcessErrors(ctx *user_context.UserContext, currentURL string) (body []byte, err error) {
	response, err := http.Get(currentURL)

	if err != nil {
		log.Fatalf("Failed to fetch response when doing $map command, user context: %v\nerr: %v\nURL: %v\n", ctx, err, currentURL)
	}

	body, err = io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to parse response when have doing $map command, user context: %v\nerr: %v\nURL: %v\n", ctx, err, currentURL)
	}
	err = response.Body.Close()
	if err != nil {
		log.Fatalf("Failed to close body of response when have received $map command, user context: %v\nerr: %v\nURL: %v\n", ctx, err, currentURL)
	}

	if response.StatusCode >= 400 {
		log.Fatalf("Response of $map failed with status code: %d\nbody: %s\n", response.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}
