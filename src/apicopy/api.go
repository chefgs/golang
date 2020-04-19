// Golang REST API program
package restapi

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)

func PostMethod(c *gin.Context) {
  fmt.Println("\napi.go 'PostMethod' called")
  message := "PostMethod called"
  c.JSON(http.StatusOK, message)
}

func GetMethod(c *gin.Context) {
  fmt.Println("\napi.go 'GetMethod' called")
  message := "GetMethod called"
  c.JSON(http.StatusOK, message)
}

func HttpFunctionCalls() {
  router := gin.Default()

  router.POST("/", PostMethod)
  router.GET("/", GetMethod)

  listenPort := "4000"
  // Listen and Server on the LocalHost:Port
  router.Run(":"+listenPort)
}
