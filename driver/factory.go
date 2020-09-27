package driver

import (
	"time"

	"github.com/ApogeeNetworking/gonetssh/client"
	"golang.org/x/crypto/ssh"
)

// Factory implements the Driver Interface
type Factory interface {
	Connect(retries int) error
	Disconnect()
	SendCmd(cmd, prompt string, delay time.Duration) (string, error)
	ReadUntil(prompt string) (string, error)
	NewClientConfig() *ssh.ClientConfig
	NewClient(cfg *ssh.ClientConfig) (*ssh.Client, error)
}

// NewDriver ...
func NewDriver(conn client.Connectioner) Factory {
	return &Driver{Connection: conn}
}
