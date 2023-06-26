package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Client struct {
	AccessToken string
	BaseUrl     string
}

func NewClient() (*Client, error) {
	token, err := GetAuthToken()
	if err != nil {
		return nil, fmt.Errorf("Can't get auth token: %w", err)
	}
	return &Client{
		AccessToken: token,
		BaseUrl:     os.Getenv("LCBOAPI_URL"),
	}, nil
}

func (c *Client) DoRequest(req *http.Request) (body []byte, err error) {
	req.Header.Set("Authorization", c.AccessToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	body, _ = ioutil.ReadAll(res.Body)
	fmt.Printf("%s\n", body)
	if res.StatusCode == 401 {
		fmt.Printf("\nToken expired.\n")
		token, err := GetAuthToken()
		if err != nil {
			return nil, fmt.Errorf("Unable to refresh auth token: %w", err)
		}
		c.AccessToken = token
		fmt.Printf("\nNew token obtained.\n")
		return c.DoRequest(req)
	}

	return body, err
}
