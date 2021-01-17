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
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(req, "--REQUEST--")

	db, err := sql.Open("mysql", "root:admin@123@tcp(127.0.0.1:3306)/amazon")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Query("Insert into products(product_url,product_name,details,image_url,price,review_count, date_updated) Values(?,?,?,?,?,?,now());",
		req.URL, req.Name, req.Details, req.ImageURL, req.Price, req.ReviewCount)
	if err != nil {
		fmt.Println("UPDATING DATA")
		_, err = db.Query("Update products set product_name = ?,details = ?, image_url =?,price =?,review_count =?,date_updated=now() where product_url = ?;",
			req.Name, req.Details, req.ImageURL, req.Price, req.ReviewCount, req.URL)
		if err != nil {
			fmt.Println(err)
		}
	}

	res := model.Response{
		Response: req.Name + "added to products table",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}
