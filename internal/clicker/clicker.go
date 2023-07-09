package clicker

import (
	"fmt"
	"io"
	"net/http"
)

type Clicker interface {
	Click(url string) ([]byte, error)
}

type clicker struct {
	client http.Client
}

func New() Clicker {
	c := http.Client{}

	return &clicker{client: c}
}

func (c *clicker) Click(url string) ([]byte, error) {
	body := []byte{}

	resp, err := c.client.Get(url)
	if err != nil {
		return body, fmt.Errorf("failed to fetch response: %s", err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return body, fmt.Errorf("failed to read response body: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return body, fmt.Errorf("response status code is not OK, body: %s", body)
	}

	return body, nil
}
