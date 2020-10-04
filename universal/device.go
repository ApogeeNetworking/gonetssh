package universal

import "golang.org/x/crypto/ssh"

// Device is the interface all Devices will implement
type Device interface {
	Connect(retries int) error
	Disconnect()
	SendCmd(cmd string) (string, error)
	SendConfig(cmds []string) (string, error)
	NewClientConfig() *ssh.ClientConfig
	NewClient(sshCfg *ssh.ClientConfig) (*ssh.Client, error)
}
