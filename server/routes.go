package server

import "github.com/gin-gonic/gin"

func setUpRouter() *gin.Engine {
  r := gin.Default()
  r.GET("/ping", ping)
  r.GET("/", homePage)
  r.GET("/about", about)
  r.GET("/resume", resume)
  r.GET("/interests", interests)
  r.GET("/art", art)
  r.GET("/donate", donate)

  return r
}

func homePage(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "Home",
    })
}

func about(c * gin.Context) {
  c.JSON(200, gin.H{
    "message": "About",
  })
}

func resume(c * gin.Context) {
  c.JSON(200, gin.H{
    "message": "Resume",
  })
}

func interests(c * gin.Context) {
  c.JSON(200, gin.H{
    "message": "Interests",
  })
}

func art(c * gin.Context) {
  c.JSON(200, gin.H{
    "message": "Art",
  })
}

func donate(c * gin.Context) {
  c.JSON(200, gin.H{
    "message": "Donate",
  })
}

func ping(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "pong",
  })
}
