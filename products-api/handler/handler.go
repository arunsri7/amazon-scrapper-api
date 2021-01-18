package handler

import (
	"amazon-api/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//AddProducts updates/inserts the product data to the db
func AddProducts(w http.ResponseWriter, r *http.Request) {
	var req model.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(req, "--REQUEST--")
	// Database Config
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		`mongodb+srv://admin:admin@cluster0.1vpl5.mongodb.net/amazon?retryWrites=true&w=majority`,
	))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	//Cancel context to avoid memory leak

	// Ping our db connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("amazon")
	productsCollection := quickstartDatabase.Collection("products")
	dateUpdated := time.Now()
	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"url", req.URL}}
	update := bson.D{
		{"$set", bson.D{{"date_updated", dateUpdated}}},
		{"$set", bson.D{{"name", req.Name}}},
		{"$set", bson.D{{"review count", req.ReviewCount}}},
		{"$set", bson.D{{"price", req.Price}}},
		{"$set", bson.D{{"details", req.Details}}},
		{"$set", bson.D{{"image", req.ImageURL}}}}
	var updatedDocument bson.M
	err = productsCollection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDocument)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		log.Fatal(err)
	}

	res := model.Response{
		Response: req.Name + "Succesfully Added to DB",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

//GetAllProducts returns all the products from the db
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	// Database Config
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		`mongodb+srv://admin:admin@cluster0.1vpl5.mongodb.net/amazon?retryWrites=true&w=majority`,
	))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	//Cancel context to avoid memory leak

	// Ping our db connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	defer client.Disconnect(ctx)
	quickstartDatabase := client.Database("amazon")
	productsCollection := quickstartDatabase.Collection("products")
	products, err := productsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var episodes []bson.M
	if err = products.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(episodes)
}

//GetProductByURL returns the object from the db given the url
func GetProductByURL(w http.ResponseWriter, r *http.Request) {
	var req model.URL
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
	}
	// Database Config
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		`mongodb+srv://admin:admin@cluster0.1vpl5.mongodb.net/amazon?retryWrites=true&w=majority`,
	))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	//Cancel context to avoid memory leak

	// Ping our db connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	defer client.Disconnect(ctx)
	quickstartDatabase := client.Database("amazon")
	productsCollection := quickstartDatabase.Collection("products")
	var product bson.M

	fmt.Println(req.URL, "response")

	if err = productsCollection.FindOne(ctx, bson.M{"url": req.URL}).Decode(&product); err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
