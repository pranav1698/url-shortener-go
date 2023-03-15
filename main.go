package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"path"
	"regexp"
	"strings"
)

// Bytes string to generate random hash
const Bytes = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var (
	links map[string]string
)

func main() {
	// creating a list of links to keep storage (in-memory) of the links and there shortened form
	links = map[string]string{}

	// adding handlers for inserting a new short link and getting the short link
	http.HandleFunc("/short", CreateShortLink)
	http.HandleFunc("/get/", Get)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

// insert: insert it into the map of links and then shorten it
// http://localhost:8080/insert?link=https://google.com
func CreateShortLink(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Query().Get("link")
	if len(link) != 0 {
		log.Println(link)

		// checks if the given link is an absolute path or not
		if !validLink(link) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Cannot create shortlink for the input, need absolute path of the url. Ex: /insert?link=https://github.com/")
			return
		}

		// check if the given link is a duplicate or not
		if !checkDuplicate(link) {
			randomString := randStringBytes(10)
			links[randomString] = link

			fmt.Println(links[link])
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusAccepted)

			shortLink := fmt.Sprintf("<a href=\"http://localhost:8080/get/%s\">http://localhost:8080/get/%s</a>", randomString, randomString)
			fmt.Fprint(w, "Shortlink inserted: ")
			fmt.Fprintln(w, shortLink)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Link already present in memory")
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Failed to add Link")
	}
}

// get: redirects to the actual link
// http://localhost:8080/get/M3Y5SqDf7A
func Get(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	log.Println("Get Link:", urlPath)

	basePath := path.Base(urlPath)
	log.Printf("Redirecting to %s", links[basePath])
	http.Redirect(w, r, links[basePath], http.StatusTemporaryRedirect)
}

// checkDuplicate will check if the link present in the links map
func checkDuplicate(link string) bool {
	for _, value := range links {
		if link == value {
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
	if r.MatchString(link) {
		return true
	}
	return false
}

// to create a random string given the length(n) of string
func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = Bytes[rand.Intn(len(Bytes))]
	}
	return string(b)
}
