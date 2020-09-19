package client

import "golang.org/x/crypto/ssh"

// Connectioner refers to the Interface Implemented by conn.SSH
type Connectioner interface {
	Connect(retries int) error
	Disconnect()
	Read() (string, error)
	Write(cmd string) int
	ExecEnable(pass string)
	NewClientConfig() *ssh.ClientConfig
	NewClient(cfg *ssh.ClientConfig) (*ssh.Client, error)
}

// NewConnection instantiate an SSHConn that implements Connection
func NewConnection(host, user, pass string) (Connectioner, error) {
	c, err := NewSSH(host, user, pass)
	if err != nil {
		return nil, err
	}
	return c, nil
}
