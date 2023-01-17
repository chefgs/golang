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

/* type urlCacheStore []struct {
	UrlToShorten  string json:"urltoshorten"
	UrlOutput string json:"urloutput"
}

func SetCache(key string, ucs interface{}) bool {
    Cache.Set(key, ucs, cache.NoExpiration)
    return true
}

func GetCache(key string) (urlCacheStore, bool) {
    var ucs urlCacheStore
    var found bool
    data, found := Cache.Get(key)
    if found {
		ucs = data.(urlCacheStore)
    }
    return ucs, found
} */

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
	} else {
		defer response.Body.Close()
		resp_body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		resp_body_str := string(resp_body)
		// Print the shortened url in console
		fmt.Println(resp_body_str)

		// Print the shortened url in browser
		c.JSON(http.StatusOK, resp_body_str)

		key := urlToCovert

		memCache.Set(key, resp_body_str, time.Minute)
		//fmt.Println(memCache.Has(key), memCache.Count())
	}
}

func GetStoredUrl(c *gin.Context) {
	//fmt.Println("\n'GetStoredUrls' called")
	urlToFetch := c.Params.ByName("urlToFetch")

	// get previously stored url
	val := memCache.Get(urlToFetch)
	fmt.Println(val)
	c.JSON(http.StatusOK, val)
}

func main() {
	router := gin.Default()

	// Get function call to shorten the url
	router.GET("/:urlToCovert", GetUrlAndCovert)

	// Create sub router to add "Get" function call to retrieve shortened url
	subRouter := router.Group("/v1/fetchcachedurl")
	subRouter.GET("/:urlToFetch", GetStoredUrl)

	listenPort := "8086"
	// Listen and Server on the LocalHost:Port
	router.Run(":" + listenPort)
}
