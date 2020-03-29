package main

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)

func postMethod(c *gin.Context) {
  fmt.Println("\napi.go 'postMethod' called")
  message := "postMethod called"
  c.JSON(http.StatusOK, message)
}

func main() {
  router := gin.Default()

  router.POST("/", postMethod)

  listenPort := "4000"
  // Listen and Server on the LocalHost:Port
  router.Run(":"+listenPort)
}
