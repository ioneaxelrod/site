package main

import "github.com/gin-gonic/gin"
import "server"

func main() {
  Run()
}

func Run() {
  server.New()
  server.Run()
}
