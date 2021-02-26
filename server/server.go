package server

import "github.com/gin-gonic/gin"

type Server struct {
  db  Database
  router *gin.Engine
}

func New() Server {
  r := gin.Default()
  	r.GET("/ping", func(c *gin.Context) {
  		c.JSON(200, gin.H{
  			"message": "pong",
  		})
  	})
  return Server{
    router: r,
    db: nil,
  }
}

func setUpRouter() *gin.Engine {
  return nil
}

func connectToDB()

(s Server) func GetDB() Database {
  return s.db
}

(s Server) func GetRouter() {
  return s.router
}

(s Server) func Run() {
  s.router.Run()
}
