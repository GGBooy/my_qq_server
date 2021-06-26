package main

import (
	"my_qq_server/server"
)

var (
	addr = ":8972"
)

func main() {
	r := server.RegisterHandle()
	defer server.CloseDB()

	r.Run(addr)
}
