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
	fmt.Println("Server Running")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/addProducts", handler.AddProducts)
	log.Fatal(http.ListenAndServe(":8000", router))
}
