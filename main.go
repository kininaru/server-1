package main

import (
	"fmt"

	"github.com/kininaru/server-1/server"
)

func main() {
	fmt.Println("Chess server listening at: 0.0.0.0:23333")
	server.InitTcpServer(23333)
}
