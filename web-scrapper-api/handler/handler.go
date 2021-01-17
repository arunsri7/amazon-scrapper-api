package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"web-scrapper/model"

	"github.com/gocolly/colly"
)

//Scrape function scrapes the data and calls the amazon-api which posts the data to db
func Scrape(w http.ResponseWriter, r *http.Request) {
	var req model.Request
	json.NewDecoder(r.Body).Decode(&req)

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", req.URL)
	})

	var (
		productName string
		ratings     string
		price       string
		details     string
		imageURL    string
	)

	c.OnHTML("div#ppd", func(e *colly.HTMLElement) {
		productName = e.ChildText("span#productTitle")
		details = e.ChildText("div#feature-bullets > ul > li:nth-child(4) > span")
		price = e.ChildText("div#olp-upd-new-used > span > a > span.a-size-base.a-color-price")
		ratings = e.ChildText("span#acrCustomerReviewText")

		fmt.Printf("Product Name: %s \n", productName)
		fmt.Printf("Product Reviews: %s \n", ratings)
		fmt.Printf("Product Price: %s \n", price)
		fmt.Printf("Product Discription: %s \n", details)
	})

	c.OnHTML("div#imgTagWrapperId", func(e *colly.HTMLElement) {
		imageURL = e.ChildAttr("img", "data-old-hires")
		fmt.Printf("Product Image: %s \n", imageURL)
	})
	c.Visit(req.URL)
	// Creating byte data to be passed as object to the POST api to save the data
	res := model.PostData{
		req.URL,
		productName,
		details,
		ratings,
		price,
		imageURL,
	}

	jsonData, err := json.Marshal(res)

	resp, err := http.Post("http://localhost:8000/addProducts",
		"application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		print(err, " error making post request")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
