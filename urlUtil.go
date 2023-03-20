package main

import "time"

// UrlUtil: structure to hold the url to be added
// Url: link to be added
// Expiration: time of expiration for the short link after which rhe link will not be valid
type UrlUtil struct {
	Url        string
	Expiration time.Time
}
