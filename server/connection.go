package server

func (this *UserConnection) tell(msg string) error {
	msg += "\n"
	_, err := this.connection.Write([]byte(msg))
	return err
}

func (this *UserConnection) read() (string, error) {
	return this.reader.ReadString('\n')
}
