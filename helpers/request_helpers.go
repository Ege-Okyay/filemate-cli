package helpers

import (
	"fmt"
	"net/http"
	"os"
)

func SendGETRequest(route string) (*http.Response, error) {
	requestURL := fmt.Sprintf(os.Getenv("API_URL") + route)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Client: Could not create request: %s\n", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Client: Error making http request: %s\n", err)
	}

	return res, nil
}
