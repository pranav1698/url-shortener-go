package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_checkDuplicateTrue(t *testing.T) {
	addLink("http://google.com", "abcd1235678")
	assert.True(t, checkDuplicate("http://google.com"))
}

func Test_checkDuplicateFalse(t *testing.T) {
	assert.False(t, checkDuplicate("http://github.com"))
}

func Test_validLinkTrue(t *testing.T) {
	valid := validLink("http://google.com")
	assert.True(t, valid)
}

func Test_validLinkFalse(t *testing.T) {
	assert.False(t, validLink("google.com"))
}

func Test_checkExpirationFalse(t *testing.T) {
	addLink("http://google.com", "abcd1235678")
	assert.False(t, checkExpiration("abcd1235678"))
}

func Test_CreateShortLink(t *testing.T) {
	req, err := http.NewRequest("GET", "/short?link=https://firefox.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateShortLink)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusAccepted, rr.Code)
}

func Test_CreateShortLinkDuplicate(t *testing.T) {
	req, err := http.NewRequest("GET", "/short?link=https://firefox.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateShortLink)
	handler.ServeHTTP(rr, req)

	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, "Link already present in memory", string(body))
}

func Test_CreateShortLinkNotValid(t *testing.T) {
	req, err := http.NewRequest("GET", "/short?link=www.google.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateShortLink)
	handler.ServeHTTP(rr, req)

	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, "Cannot create shortlink for the input, need absolute path of the url. Ex: /short?link=https://github.com/", string(body))
}

func Test_CreateShortLinkEmpty(t *testing.T) {
	req, err := http.NewRequest("GET", "/short", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateShortLink)
	handler.ServeHTTP(rr, req)

	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, "Link to short not found, Try: http://localhost:8080/short?link=https://google.com", string(body))
}

func Test_Get(t *testing.T) {
	addLink("https://youtube.com", "09iopcnjnc")
	req, err := http.NewRequest("GET", "/get/09iopcnjnc", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Get)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusTemporaryRedirect, rr.Code)
}
