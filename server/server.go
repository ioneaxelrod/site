package server

import (
  "github.com/gin-gonic/gin"
)

// Server is the powerhouse of the cell
type Server struct {
  db  interface{}
  router *gin.Engine
}

// New creates new Server for people
func New() Server {
  return Server{
    router: setUpRouter(),
    db: connectToDB(),
  }
}

// GetRouter rules
func (s Server) GetRouter() *gin.Engine {
  return s.router
}

// Run gets the server running
func (s Server) Run() {
  s.router.Run()
}

func setUpRouter() *gin.Engine {
  r := gin.Default()
  r.GET("/ping", ping)
  r.GET("/", showHomePage)
  return r
}

func connectToDB() interface{} {
  return nil
}
