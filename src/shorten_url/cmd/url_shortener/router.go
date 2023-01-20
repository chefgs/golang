package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/liyaojian/cache"
)

// var memCache = cache.New(5*time.Minute, 5*time.Minute)
var memCache = cache.NewFileCache("./cache")

func GetUrlAndCovert(c *gin.Context) {
	//fmt.Println("\n'GetUrlAndCovert' called")
	urlToCovert := c.Params.ByName("urlToCovert")
	//message := "GetMethod Called With Param: " + urlValue

	// Append the Url to tinyurl API to shorten
	baseUrl := "http://tinyurl.com/api-create.php?url="
	getReqUrl := baseUrl + urlToCovert

	// Execute the API call to tinyurl API to shorten
	response, err := http.Get(getReqUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	resp_body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	resp_body_str := string(resp_body)
	// Print the shortened url in console
	fmt.Println(resp_body_str)

	// Save the converted URL data in Cache file for fetching it later
	// Key to store the data will be URL we are converting, and converted shorted URL will be stored as data
	key := urlToCovert
	memCache.Set(key, resp_body_str, time.Minute)
	//fmt.Println(memCache.Has(key), memCache.Count())

	// Print the shortened url in browser if converted successfully
	c.JSON(http.StatusOK, resp_body_str)
}

func GetStoredUrl(c *gin.Context) {
	//fmt.Println("\n'GetStoredUrls' called")
	urlToFetch := c.Params.ByName("urlToFetch")

	// get previously stored url
	val := memCache.Get(urlToFetch)
	fmt.Println(val)
	c.JSON(http.StatusOK, val)
}

func setup_router() *gin.Engine {
	gin_router := gin.Default()

	// Create sub router to version api
	subRouter := gin_router.Group("/v1")

	// Get function call to shorten the url
	// Creating a GET request to test the API in browser address bar.
	// Otherwise this is supposed to be a POST request
	subRouter.GET("/:urlToCovert", GetUrlAndCovert)

	// Add "Get" function call to retrieve shortened url
	subRouter.GET("/fetchcachedurl/:urlToFetch", GetStoredUrl)

	return gin_router
}
