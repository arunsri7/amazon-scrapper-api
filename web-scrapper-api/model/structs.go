package model

type Request struct {
	URL string
}

//Request object
type PostData struct {
	URL         string
	Name        string
	Details     string
	ReviewCount string
	Price       string
	ImageURL    string
}
