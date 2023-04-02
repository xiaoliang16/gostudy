package main

import (
	"gostudy/conf"
	"gostudy/server"
)

func main() {
	conf.Init()
	r := server.NewRouter()
	r.Run(":3000")
}