package main

import (
	"fmt"

	"github.com/kininaru/server-1/server"
)

func main() {
	fmt.Println("server is gonna run in port 23333")
	server.InitTcpServer(23333)
}
