package main

import "github.com/ioneaxelrod/site/server"

func main() {
  run()
}

func run() {
  s := server.New()
  s.Run()
}
