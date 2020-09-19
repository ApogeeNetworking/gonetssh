package driver

import (
	"time"

	"github.com/drkchiloll/gonetmiko/conn"
	"golang.org/x/crypto/ssh"
)

// Factory implements the Driver Interface
type Factory interface {
	Connect(retries int) error
	Disconnect()
	SendCmd(cmd, regex string, dur time.Duration) (string, error)
	ReadUntil(regex string) (string, error)
	ExecEnable(enablePass string)
	NewClientConfig() *ssh.ClientConfig
	NewClient(cfg *ssh.ClientConfig) (*ssh.Client, error)
}

// NewDriver ...
func NewDriver(conn conn.Connectioner) Factory {
	return &Driver{Connection: conn}
}
