package main

import (
	"fmt"
	"log"
	"net/http"
	"web-scrapper/handler"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Amaozon Api server is running")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/scrapeAmazonData", handler.Scrape)
	log.Fatal(http.ListenAndServe(":5001", router))
}
