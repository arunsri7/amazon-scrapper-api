package main

import (
	"amazon-api/handler"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Amazon-product-api is running on port 8000")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/addProducts", handler.AddProducts)
	router.HandleFunc("/getAllProducts", handler.GetAllProducts)
	router.HandleFunc("/getProductByURL", handler.GetProductByURL)
	log.Fatal(http.ListenAndServe(":8000", router))
}
