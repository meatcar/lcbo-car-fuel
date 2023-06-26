package client

import (
	"net/http"
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatalf(`NewClient() = %q, %v, want Client{}, error`, client, err)
	}
}

func TestDoRequest(t *testing.T) {
	client := Client{
		AccessToken: `eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IjJaUXBKM1VwYmpBWVhZR2FYRUpsOGxWMFRPSSIsImtpZCI6IjJaUXBKM1VwYmpBWVhZR2FYRUpsOGxWMFRPSSJ9.eyJhdWQiOiJhcGk6Ly8wZDFkOGYwOS02MjgzLTQ3YzYtODFkMC05MWM0OGYwNDU3NGIiLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC80YTViNzk0Mi1iMGIxLTQyNDQtYjYwMS0xOWY4ZTZjMzNlNDgvIiwiaWF0IjoxNjU3MzIzNjY3LCJuYmYiOjE2NTczMjM2NjcsImV4cCI6MTY1NzMyNzU2NywiYWlvIjoiRTJaZ1lIQXZVZU9mdk90RCsxN2hwSlU3cDY1YkFnQT0iLCJhcHBpZCI6IjY0M2I5ZTdhLTRmYzMtNGY4Zi1hYTBkLTBjZmM2N2U5N2E3MCIsImFwcGlkYWNyIjoiMSIsImlkcCI6Imh0dHBzOi8vc3RzLndpbmRvd3MubmV0LzRhNWI3OTQyLWIwYjEtNDI0NC1iNjAxLTE5ZjhlNmMzM2U0OC8iLCJvaWQiOiIzYzM1NDZlYS01NDhmLTRiY2ItOWU2MC03ZjllN2QxYmExYzIiLCJyaCI6IjAuQVRnQVFubGJTckd3UkVLMkFSbjQ1c00tU0FtUEhRMkRZc1pIZ2RDUnhJOEVWMHM0QUFBLiIsInN1YiI6IjNjMzU0NmVhLTU0OGYtNGJjYi05ZTYwLTdmOWU3ZDFiYTFjMiIsInRpZCI6IjRhNWI3OTQyLWIwYjEtNDI0NC1iNjAxLTE5ZjhlNmMzM2U0OCIsInV0aSI6InZrZlVnXzYzSVV1Umc2TU9neWt6QUEiLCJ2ZXIiOiIxLjAifQ.OWVY1k44TW72KVKM3qe-427U7ULe5udNHb0mIuK4jsk1SgLDjPrYgVdBogVAKGYOfNHm7yKOTdo1QeXTleM1fMmUK10x3UG1wcjJLbWPoJUyTXLH9ROP5w9PZ3H6mVRyWE9WIxKI6qClbAJvZAgHQo1PZ_jq8soD_JvkaQjLDgqYflsLtwsB6UuqZq1A_mmh_75m0iNmrZDTerSpbvdYqvejg8i5pC0XxzJY53omjkgko63_b4BgRdugR45wkoISj7Urp-kHJsOXlhoazbSA8UNintOO5uw_q-dO6Bc7jstgNsC4e75gOKfV1BPtQJyKWty7gL33r_HIqWiEQINDIQ`,
	}
	req, _ := http.NewRequest("GET", os.Getenv("LCBOAPI_URL"), nil)
	data, err := client.DoRequest(req)
	if err != nil || data == nil {
		t.Fatalf(`DoRequest() = %q, %v, want []byte, error`, data, err)
	}
}
