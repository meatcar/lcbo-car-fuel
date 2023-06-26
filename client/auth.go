package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

func GetAuthToken() (out string, err error) {
	data := url.Values{}
	data.Add("grant_type", "client_credentials")
	data.Add("client_id", os.Getenv("LCBOAPI_CLIENT_ID"))
	data.Add("client_secret", os.Getenv("LCBOAPI_CLIENT_SECRET"))
	data.Add("scope", os.Getenv("LCBOAPI_SCOPE"))

	tennant := os.Getenv("LCBOAPI_TENNANT")
	baseurl := "https://login.microsoftonline.com/%s/oauth2/v2.0/token"

	res, err := http.PostForm(fmt.Sprintf(baseurl, tennant), data)
	if err != nil {
		return
	}
	if res.StatusCode != 200 {
		err = errors.New(res.Status)
		return
	}

	var token AuthResponse
	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		return
	}
	out = token.AccessToken
	return
}
