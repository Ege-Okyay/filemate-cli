package helpers

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func SendHttpRequest(route string, method string, jsonBody []byte, multipartBody *bytes.Buffer, multipartBoundary string, authToken string) (*http.Response, error) {
	requestURL := fmt.Sprintf(os.Getenv("API_URL") + route)

	var req *http.Request
	var reqErr error

	switch method {
	case "GET":
		req, reqErr = http.NewRequest(http.MethodGet, requestURL, nil)

		break
	case "POST":
		if multipartBody == nil {
			// The request is sending a json content
			req, reqErr = http.NewRequest(http.MethodPost, requestURL, bytes.NewReader(jsonBody))

			req.Header.Set("Content-Type", "application/json")
		} else {
			// The request is sending a file
			req, reqErr = http.NewRequest(http.MethodPost, requestURL, bytes.NewReader(multipartBody.Bytes()))

			if multipartBoundary == "" {
				reqErr = fmt.Errorf("Empty multipart/form-data boundary")
				break
			}

			contentTypeHeader := fmt.Sprintf("multipart/form-data; boundary=%s", multipartBoundary)

			req.Header.Set("Content-Type", contentTypeHeader)
		}

		break
	default:
		return nil, fmt.Errorf("Client: Unsupported HTTP method: %s\n", method)
	}

	if authToken != "" {
		req.Header.Set("Authorization", authToken)
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
