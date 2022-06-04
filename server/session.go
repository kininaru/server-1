package server

import (
	"bufio"
	"net"
)

type UserConnection struct {
	connection net.Conn
	buf        []byte
	scanner    bufio.Scanner
}

func session(connection net.Conn) {
	userConnection := &UserConnection{
		connection: connection,
		scanner:    *bufio.NewScanner(connection),
	}
	if err := userConnection.login(); err != nil {
		userConnection.closeSession()
		return
	}

	userConnection.closeSession()
}

func (this *UserConnection) login() error {

}

func (this *UserConnection) closeSession() {
	this.tell("end")
}
