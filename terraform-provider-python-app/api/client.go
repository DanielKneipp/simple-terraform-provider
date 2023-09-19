package api

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Client struct {
	endpoint   string
	httpClient *http.Client
}

func NewClient(endpoint string) *Client {
	return &Client{
		endpoint:   endpoint,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetAll() ([]int, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("%s/all", c.endpoint))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
	}

	body_string := string(body_bytes)
	if body_string == "" {
		return make([]int, 0), nil
	}

	numbers_str := strings.Split(body_string, ",")

	numbers := make([]int, len(numbers_str))
	for i, s := range numbers_str {
		numbers[i], err = strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
	}

	return numbers, nil
}

func (c *Client) AddNumber(number int) error {
	resp, err := c.httpClient.Post(fmt.Sprintf("%s/%d", c.endpoint, number), "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
	}

	return nil
}

func (c *Client) RemoveAll() error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/all", c.endpoint), nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
	}

	return nil
}
