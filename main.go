package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/meatcar/lcbo-car-fuel/client"
)

func GetProducts(c *client.Client, page int, limit int) (response client.ProductResponse, err error) {
	req, err := c.ProductsRequest()
	if err != nil {
		return
	}
	client.PaginateRequest(req, page, limit)
	body, err := c.DoRequest(req)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &response)
	return
}

func main() {
	ts := time.Now()

	path := fmt.Sprintf("data/%04d-%02d-%02dT%02d:%02d:%02d",
		ts.Year(), ts.Month(), ts.Day(), ts.Hour(), ts.Minute(), ts.Second())
	err := os.Mkdir(path, 0750)
	if err != nil {
		log.Fatalf("Unable to make dir %s", path)
	}

	path = fmt.Sprintf("%s/products.json", path)
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("Unable to open file %s", path)
	}

	c, err := client.NewClient()
	if err != nil {
		log.Fatalf("Error starting client: %v", err)
	}

	res, err := GetProducts(c, 1, 1)
	if err != nil {
		log.Fatalf("Error getting initial products: %v", err)
	}

	start := time.Now()
	page, limit := 25, 200
	count := res.ResultCount
	for (page * limit) < count {
		res, err := GetProducts(c, page, limit)
		if err != nil {
			log.Fatalf("Error getting products: %v", err)
		}

		page += 1

		for i, product := range res.Products {
			cur := (page-1)*limit + i + 1
			elapsed := time.Now().Sub(start)
			avg := elapsed.Seconds() / float64(cur)
			duration := fmt.Sprintf("%.fs", float64(count-cur)*avg)
			remaining, _ := time.ParseDuration(duration)
			if err != nil {
				log.Fatalf("Unable to parse duration %s: %v", duration, err)
			}
			fmt.Printf("\rFetching %d/%d %s elapsed (%.4f avg) %s remaining ...\033[K", cur, count, elapsed, avg, remaining)

			req, err := c.ProductRequest(product.ItemNumber)
			if err != nil {
				log.Fatalf("Error creating product request: %v", err)
			}

			body, err := c.DoRequest(req)
			if err != nil {
				log.Fatalf("Error fetching product %q: %v", product, err)
			}

			_, err = file.Write(body)
			if err != nil {
				log.Fatalf("Error writing to file: %v", err)
			}
		}
	}
}
