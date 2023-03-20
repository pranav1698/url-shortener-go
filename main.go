package main

import (
	"log"
	"net/http"
)

// LinksStore: storage for our links and their shortened versions
// Fields:
// Links: map to store the urls and their shortened links
// MaximumUrls: maximum links that we can store
// CurrentUrls: current links present in the store
type LinkStore struct {
	Links       map[UrlUtil]string
	MaximumUrls int
	CurrentUrls int
}

// initializing our LinksStore
var LinksStore = LinkStore{
	Links:       map[UrlUtil]string{},
	MaximumUrls: 10,
	CurrentUrls: 0,
}

func main() {
	// adding handlers for inserting a new short link and getting the short link
	http.HandleFunc("/short", CreateShortLink)
	http.HandleFunc("/get/", Get)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
