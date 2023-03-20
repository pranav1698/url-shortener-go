package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"path"
	"regexp"
	"strings"
	"time"
)

// Bytes string to generate random hash
const Bytes = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// short: insert it into the map of Links and then shorten it
// http://localhost:8080/short?link=https://google.com
func CreateShortLink(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Query().Get("link")
	if len(link) != 0 {
		log.Println(link)

		// checks if the given link is an absolute path or not
		if !validLink(link) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Cannot create shortlink for the input, need absolute path of the url. Ex: /short?link=https://github.com/")
			return
		}

		// checking if the number of short links is less than maximum links that we can store
		if LinksStore.CurrentUrls >= LinksStore.MaximumUrls {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Maximum url limit exceeded, please wait for some time....")
			return
		}

		// check if the given link is a duplicate or not
		if checkDuplicate(link) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Link already present in memory")
			return
		}

		randomString := randStringBytes(10)
		addLink(link, randomString)

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusAccepted)

		shortLink := fmt.Sprintf("<a href=\"http://localhost:8080/get/%s\">http://localhost:8080/get/%s</a>", randomString, randomString)
		fmt.Fprint(w, "Shortlink inserted: ")
		fmt.Fprintln(w, shortLink)
		return

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Link to short not found, Try: http://localhost:8080/short?link=https://google.com")
	}
}

// get: redirects to the actual link
// http://localhost:8080/get/M3Y5SqDf7A
func Get(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	log.Println("Get Link:", urlPath)

	basePath := path.Base(urlPath)
	for urlUtil, shortLink := range LinksStore.Links {
		if basePath == shortLink {
			if !checkExpiration(shortLink) {
				log.Printf("Redirecting to %s", &urlUtil)
				http.Redirect(w, r, urlUtil.Url, http.StatusTemporaryRedirect)
			} else {
				delete(LinksStore.Links, urlUtil)
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintln(w, "Link Expired, create new one...")
			}
		}
	}
}

// Adding link to our link store
func addLink(link string, randomString string) string {
	urlUtil := UrlUtil{
		Url:        link,
		Expiration: time.Now().Add(time.Minute * 2),
	}

	LinksStore.Links[urlUtil] = randomString
	LinksStore.CurrentUrls++

	return randomString
}

// checkDuplicate will check if the link present in the Links map
func checkDuplicate(link string) bool {
	for urlUtil, _ := range LinksStore.Links {
		if urlUtil.Url == link {
			return true
		}
	}

	return false
}

// checks if the link is an absolute path or not
// http.Redirects requires an absolute path to reach a link outside the application
func validLink(link string) bool {
	r, err := regexp.Compile("^(http|https)://")
	if err != nil {
		return false
	}
	link = strings.TrimSpace(link)
	log.Printf("Checking for valid link: %s", link)

	// Check if string matches the regex
	return r.MatchString(link)
}

// To create a random string given the length(n) of string
func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = Bytes[rand.Intn(len(Bytes))]
	}
	return string(b)
}

// Checking if the short link is expired or not
func checkExpiration(link string) bool {
	for urlUtil, shortLink := range LinksStore.Links {
		if shortLink == link {
			if time.Now().After(urlUtil.Expiration) {
				return true
			}
		}
	}

	return false
}
