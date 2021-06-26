package main

import (
	"my_qq_server/server"
)

var (
	addr = ":8972"
)

func main() {
	r := server.RegisterHandle()

	r.Run(addr)
}
