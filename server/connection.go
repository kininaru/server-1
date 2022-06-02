package server

import "errors"

func (this *User) tell(msg string) error {
	msg += "\n"
	_, err := this.connection.Write([]byte(msg))
	return err
}

func (this *User) read() (string, error) {
	var err error
	if this.scanner.Scan() {
		err = errors.New("no text")
	}
	return this.scanner.Text(), err
}
