package handler

import (
	"amazon-api/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AddProducts(w http.ResponseWriter, r *http.Request) {
	var req model.Request
	json.NewDecoder(r.Body).Decode(&req)
	db, err := sql.Open("mysql", "root:admin@123@tcp(127.0.0.1:3306)/amazon")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Query("Insert into products(product_url,product_name,details,image_url,price,review_count, date_updated) Values(?,?,?,?,?,?,now());",
		req.URL, req.Product.Name, req.Product.Details, req.Product.ImageURL, req.Product.Price, req.Product.ReviewCount)
	if err != nil {
		fmt.Println(err)
	}

	res := model.Response{
		Response: req.Product.Name + " added to products table",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
