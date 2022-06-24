package server

import (
	"bufio"
	"net"
	"time"
)

type UserConnection struct {
	connection net.Conn
	reader     bufio.Reader
	sid        int
}

func session(connection net.Conn) {
	userConnection := &UserConnection{
		connection: connection,
		reader:     *bufio.NewReader(connection),
	}
	userConnection.tell("y")

	userConnection.sid = len(UserList)
	UserList = append(UserList, userConnection)

	targetSid := -1
	if userConnection.sid%2 == 0 {
		for userConnection.sid == len(UserList)-1 {
			time.Sleep(10 * time.Second)
		}
		targetSid = userConnection.sid + 1
		userConnection.tell("0")
	} else {
		userConnection.tell("1")
		targetSid = userConnection.sid - 1
	}

	var msg string
	var err error
	for {
		msg, err = userConnection.read()
		if err != nil {
			UserList[targetSid].tell("ed")
			break
		}
		err := UserList[targetSid].tell(msg)
		if msg == "ed" || err != nil {
			break
		}
	}

	userConnection.tell("ed")
}
