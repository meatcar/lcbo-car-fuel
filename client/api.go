package client

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) MunicipalitiesRequest() (req *http.Request, err error) {
	url := fmt.Sprintf("%s/municipality/", c.BaseUrl)
	return http.NewRequest("GET", url, nil)
}

func (c *Client) StoresRequest(municipality string) (req *http.Request, err error) {
	url, err := url.Parse(fmt.Sprintf("%s/stores/", c.BaseUrl))
	if err != nil {
		return
	}

	q := url.Query()
	q.Set("municipality", municipality)
	url.RawQuery = q.Encode()

	return http.NewRequest("GET", fmt.Sprintf("%s", url), nil)
}

func (c *Client) ProductsRequest() (req *http.Request, err error) {
	url := fmt.Sprintf("%s/products/", c.BaseUrl)
	return http.NewRequest("GET", url, nil)
}

func (c *Client) ProductRequest(id string) (req *http.Request, err error) {
	url := fmt.Sprintf("%s/products/%s/", c.BaseUrl, id)
	return http.NewRequest("GET", url, nil)
}

func PaginateRequest(req *http.Request, page int, limit int) {
	q := req.URL.Query()
	q.Set("paginationOffset", fmt.Sprintf("%d", page))
	q.Set("numProductsPerPage", fmt.Sprintf("%d", limit))
	req.URL.RawQuery = q.Encode()
}
