package client

import (
	"testing"
)

func TestGetAuthToken(t *testing.T) {
	token, err := GetAuthToken()
	if err != nil || token == "" {
		t.Fatalf(`GetAuthToken() = %q, %v, want "<token>", error`, token, err)
	}
}
