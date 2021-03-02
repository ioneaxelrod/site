package server

import "github.com/gin-gonic/gin"

func setUpRouter() *gin.Engine {
  r := gin.Default()
  r.GET("/ping", ping)
  r.GET("/", homePage)
  return r
}

func homePage(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "poop",
    })
}

// Ping is a test endpoint
func ping(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "pong",
  })
}
