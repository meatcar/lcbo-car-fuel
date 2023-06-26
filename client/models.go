package client

type Municipality string

type Store struct {
	LocationNumber int    `json:"locationNumber"`
	LocationName   string `json:"locationName"`
}

type Product struct {
	ItemNumber string `json:"itemNumber"`
	ItemName   string `json:"itemName"`
}

type ProductResponse struct {
	Products    []Product `json:"products"`
	ResultCount int       `json:"resultCount"`
}
