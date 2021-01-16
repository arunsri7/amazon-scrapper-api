package model

import (
	"time"
)

//Request object
type Request struct {
	URL     string
	Product struct {
		Name        string
		Details     string
		ReviewCount int
		Price       int
		ImageURL    string
		time        time.Time
	}
}

//Response object
type Response struct {
	Response string
}
