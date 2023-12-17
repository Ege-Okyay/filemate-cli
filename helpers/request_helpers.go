package helpers

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

// SendHttpRequest sends an HTTP request to the specified route using the given method.
// It supports GET and POST methods with optional JSON or multipart/form-data bodies.
// It returns the HTTP response or an error if any occurs during the request.
func SendHttpRequest(route string, method string, jsonBody []byte, multipartBody *bytes.Buffer, multipartBoundary string, authToken string) (*http.Response, error) {
	// Construct the full request URL
	requestURL := fmt.Sprintf(os.Getenv("API_URL") + route)

	var req *http.Request
	var reqErr error

	// Create an HTTP request based on the specified method
	switch method {
	case "GET":
		// Create a GET request with no body
		req, reqErr = http.NewRequest(http.MethodGet, requestURL, nil)
		break
	case "POST":
		if multipartBody == nil {
			// Create a POST request with a JSON body
			req, reqErr = http.NewRequest(http.MethodPost, requestURL, bytes.NewReader(jsonBody))

			// Set the Content-Type header for JSON
			req.Header.Set("Content-Type", "application/json")
		} else {
			// Create a POST request with a multipart/form-data body
			req, reqErr = http.NewRequest(http.MethodPost, requestURL, bytes.NewReader(multipartBody.Bytes()))

			// Check if the multipartBoundary is provided
			if multipartBoundary == "" {
				reqErr = fmt.Errorf("Empty multipart/form-data boundary")
				break
			}

			// Set the Content-Type header for multipart/form-data with the specified boundary
			contentTypeHeader := fmt.Sprintf("multipart/form-data; boundary=%s", multipartBoundary)
			req.Header.Set("Content-Type", contentTypeHeader)
		}
		break
	default:
		// Unsupported HTTP method
		return nil, fmt.Errorf("Client: Unsupported HTTP method: %s\n", method)
	}

	// Set the Authorization header if an authToken is provided
	if authToken != "" {
		req.Header.Set("Authorization", authToken)
	}

	// Check if there was an error creating the request
	if reqErr != nil {
		return nil, fmt.Errorf("Client: Could not create request: %s\n", reqErr)
	}

	// Send the HTTP request and return the response or an error
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Client: Error making http request: %s\n", err)
	}

	return res, nil
}
