package client

import (
	"testing"
)

func setupTest(t *testing.T) *Client {
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	return client
}

func TestFetchMunicipalities(t *testing.T) {
	client := setupTest(t)
	data, err := client.FetchMunicipalities()
	if err != nil || data == nil {
		t.Fatalf(`FetchMunicipalities() = %q, %v, want [...], error`, data, err)
	}
}

func TestFetchStores(t *testing.T) {
	client := setupTest(t)
	data, err := client.FetchStores("TORONTO-CENTER")
	if err != nil || data == nil {
		t.Fatalf(`FetchStores() = %q, %v, want [...], error`, data, err)
	}
}

func TestFetchProducts(t *testing.T) {
	client := setupTest(t)
	data, err := client.FetchProducts()
	t.Fatalf("%s", data)
	if err != nil || data == nil {
		t.Fatalf(`FetchStores() = %q, %v, want [...], error`, data, err)
	}
}
