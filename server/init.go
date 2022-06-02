package server

import (
	"fmt"
	"net"
)

func InitTcpServer(port int) {
	server, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	var connection net.Conn
	for err == nil {
		connection, err = server.Accept()
		go session(connection)
	}
	panic(err)
}
