package conn

// Connectioner refers to the Interface Implemented by conn.SSH
type Connectioner interface {
	Connect(retries int) error
	Disconnect()
	Read() (string, error)
	Write(cmd string) int
	ExecEnable(pass string)
}

// NewConnection instantiate an SSHConn that implements Connection
func NewConnection(host, user, pass string) (Connectioner, error) {
	c, err := NewSSH(host, user, pass)
	if err != nil {
		return nil, err
	}
	return c, nil
}
