// Golang REST API unit testing program
package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetUrlAndCovertMethod(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	body := bytes.NewBuffer([]byte("{\"UrlShortenerApiTest\":\"GetUrlAndCovert\"}"))
	req, err := http.NewRequest("GET", "/v1/apitest.com", body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Setup your router, just like you did in your main function, and
	// register your routes
	router := gin.Default()
	router.GET("/v1/:urlToCovert", GetUrlAndCovert)

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	/*
		 	// Check to see if the response was what you expected
			if w.Code == http.StatusOK {
				fmt.Println(w.Body)
				t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
			} else {
				t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
			}
	*/
}

func TestGetStoredUrlMethod(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	body := bytes.NewBuffer([]byte("{\"UrlShortenerApiTest\":\"GetStoredUrl\"}"))
	req, err := http.NewRequest("GET", "/v1/fetchcachedurl/apitest.com", body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Setup your router, just like you did in your main function, and
	// register your routes
	router := gin.Default()
	router.GET("/v1/fetchcachedurl/:urlToFetch", GetStoredUrl)

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	/*
		 	// Check to see if the response was what you expected
			if w.Code == http.StatusOK {
				fmt.Println(w.Body)
				t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
			} else {
				t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
			}
	*/
}
