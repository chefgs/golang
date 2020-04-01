package main

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)

func PostMethod(c *gin.Context) {
  fmt.Println("\napi.go 'postMethod' called")
  message := "PostMethod called"
  c.JSON(http.StatusOK, message)
}

func main() {
  router := gin.Default()

  router.POST("/", PostMethod)

  listenPort := "4000"
  // Listen and Server on the LocalHost:Port
  router.Run(":"+listenPort)
}
