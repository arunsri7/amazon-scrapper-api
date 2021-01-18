package main

import (
	"fmt"
	"log"
	"net/http"
	"web-scrapper/handler"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Web Scrapper Api server is running on port :5002")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/scrapeAmazonData", handler.Scrape)
	log.Fatal(http.ListenAndServe(":5002", router))
}
