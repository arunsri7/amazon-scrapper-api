# amazon-scrapper-api

Amazon-scrapper-api(web-scapper-api) scrapes the data from the given amazon product page and calls another api(product-api) which pushes the scrapped data to mongo db atlas.

## Directory tree

	golang-api
	├── README.md
	├── docker-compose.yml
	├── products-api
	│   ├── Dockerfile
	│   ├── db.sql
	│   ├── go.mod
	│   ├── go.sum
	│   ├── handler
	│   │   └── handler.go
	│   ├── main.go
	│   └── model
	│       └── structs.go
	└── web-scrapper-api
	    ├── Dockerfile
	    ├── go.mod
	    ├── go.sum
	    ├── handler
	    │   └── handler.go
	    ├── main.go
	    └── model
		└── structs.go

## Setup

	* Git clone git@github.com:arunsri7/amazon-scrapper-api.git
	* cd amazon-scrapper-api
	* docker-compose up --build -d
	* docker-compose up
	
## API Endpoints :
  * web-scrapper-api
	  * Has one endpoint : scrapeAmazonData
    * Sample request: 
	  * URL : localhost:5002/scrapeAmazonData
	  * Method: Post
	  * Sample request body: ``` { "URL"   :  "https://www.amazon.com/Dual-Band-WiFi-Range-Extender/dp/B08DHCC8BN/ref=pd_vtp_2?pd_rd_w=H4bar&pf_rd_p=4f2ab3e8-468a-4a7c-9b91-89d6a9221c29&pf_rd_r=WH3JJ8KYN2P5NKC5GB3Y&pd_rd_r=36492927-2219-4678-b4f6-a69e0f1d1a67&pd_rd_wg=kRVbo&pd_rd_i=B08DHCC8BN&th=1"}```

  * products-api has 3 endpoints 
    * addProducts
      * endpoint :localhost:8000/addProducts
      * Method: Post
      * Sample request body: 
        {
        "URL"   :  "",
      	"Name"    :  "",
     	"Details" :   "",
      	"ReviewCount": "",
      	"Price"      : "",
     	"ImageURL"   : ""
      }
      * Sample response:
      ``` {"Successfully added to db"} ```


    * getAllProducts
      * endpoint : localhost:8000/getAllProducts
      * Method: GET
      * Sample response:
        ``` [{
            "Details": "nice crunchy",
            "_id": "600546efbb508294799d2263",
            "date_updated": "2021-01-18T14:07:11.311+05:30",
            "image": "photod.google.com",
            "name": "dateUpdated",
            "price": "250",
            "reviews": "1",
            "url": "amazon.com"},....] 


    * getProductById
        * endpoint : localhost:8000/getAllProducts
        * Method: GET
        * Sample request body: 
          {
              "URL"   :  ""
          }
        * Sample response:
         ``` {
              "Details": "nice crunchy",
              "_id": "600546efbb508294799d2263",
              "date_updated": "2021-01-18T14:07:11.311+05:30",
              "image": "photod.google.com",
              "name": "dateUpdated",
              "price": "250",
              "reviews": "1",
              "url": "amazon.com"} 
	

## Limitations
	* For some reason Amazon doens't return consistent data for scrapping. Right now even if the scrapper doesn't 
	get the required data it pushes the url to the db. If after it gets the data after retries, it updates the 
	respective data in the db. 

## TODO

	* Use ENV variables
	* Appropriate naming for variables, api, etc
	* Use a response for the web-scrapper-api
	* Add better comments
