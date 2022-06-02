package server

import (
	"bufio"
	"net"
)

type User struct {
	connection net.Conn
	buf        []byte
	scanner    bufio.Scanner
}

func session(connection net.Conn) {
	user := &User{
		connection: connection,
		scanner:    *bufio.NewScanner(connection),
	}
	if err := user.login(); err != nil {
		user.closeSession()
		return
	}

	user.closeSession()
}

func (this *User) login() error {

}

func (this *User) closeSession() {
	this.tell("end")
}
