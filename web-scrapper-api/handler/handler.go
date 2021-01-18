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
		fmt.Printf("Review Count: %s \n", ratings)
		fmt.Printf("Product Price: %s \n", price)
		fmt.Printf("Product Details: %s \n", details)
	})

	c.OnHTML("div#imgTagWrapperId", func(e *colly.HTMLElement) {
		imageURL = e.ChildAttr("img", "data-old-hires")
		fmt.Printf("Image URL : %s \n", imageURL)
	})
	c.Visit(req.URL)
	// Creating byte data to be passed as object to the POST api to save the data
	postData := model.PostData{
		req.URL,
		productName,
		details,
		ratings,
		price,
		imageURL,
	}

	jsonData, err := json.Marshal(postData)

	_, err = http.Post("http://172.31.1.3:8000/addProducts",
		"application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		print(err, " error making post request")
	}
	res := model.Response{
		Response: "Sucessfully added to the db",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}
