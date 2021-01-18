package model

import (
	"time"
)

//Request object
type Request struct {
	URL         string
	Name        string
	Details     string
	ReviewCount string
	Price       string
	ImageURL    string
	time        time.Time
}

//Response object
type Response struct {
	Response string
}

//URL object
type URL struct {
	URL string
}
