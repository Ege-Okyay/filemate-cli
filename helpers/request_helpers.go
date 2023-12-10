package helpers

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func SendHttpRequest(route string, method string, jsonBody []byte) (*http.Response, error) {
	requestURL := fmt.Sprintf(os.Getenv("API_URL") + route)

	var req *http.Request
	var reqErr error

	switch method {
	case "GET":
		req, reqErr = http.NewRequest(http.MethodGet, requestURL, nil)
		break
	case "POST":
		if jsonBody == nil {
			return nil, fmt.Errorf("Client: Empty request body!")
		}

		bodyReader := bytes.NewReader(jsonBody)
		req, reqErr = http.NewRequest(http.MethodPost, requestURL, bodyReader)

		req.Header.Set("Content-Type", "application/json")
		break
	default:
		break
	}

	if reqErr != nil {
		return nil, fmt.Errorf("Client: Could not create request: %s\n", reqErr)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Client: Error making http request: %s\n", err)
	}

	return res, nil
}
