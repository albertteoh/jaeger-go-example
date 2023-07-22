package xhttp

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	statusNotOKErr = errors.New("error response on HTTP request")
)

// Do executes an HTTP request and returns the response body as a string.
// Non-200 response codes will be returned as an error with the response body.
func Do(req *http.Request) (string, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed http request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("StatusCode: %d, Body: %s: %w", resp.StatusCode, body, statusNotOKErr)
	}
	return string(body), nil
}
