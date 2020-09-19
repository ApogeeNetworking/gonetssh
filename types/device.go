package types

import "golang.org/x/crypto/ssh"

// Device ...
type Device interface {
	Connect(retries int) error
	Disconnect()
	SendCmd(cmd string) (string, error)
	NewClientConfig() *ssh.ClientConfig
	NewClient() (*ssh.Client, error)
}
