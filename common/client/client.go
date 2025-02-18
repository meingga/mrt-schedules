package client

import (
	"fmt"
	"io"
	"net/http"
)

func DoRequest(client *http.Client, url string) ([]byte, error) {
	resp, err := client.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
