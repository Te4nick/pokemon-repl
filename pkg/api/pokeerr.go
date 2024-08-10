package api

import (
	"fmt"
)

type HTTPStatusError struct {
	StatusCode int
	URL        string
}

func (e HTTPStatusError) Error() string {
	return fmt.Sprintf("HTTP %s Status Code %d ", e.URL, e.StatusCode)
}
