package server

import "github.com/gin-gonic/gin"

func setUpRouter() *gin.Engine {
  r := gin.Default()
  r.GET("/ping", ping)
  r.GET("/", showHomePage)
  return r
}

func showHomePage(c *gin.Context) {
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
