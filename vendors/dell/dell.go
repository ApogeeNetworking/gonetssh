package dell

import (
	"github.com/ApogeeNetworking/gonetssh/driver"
	"github.com/ApogeeNetworking/gonetssh/universal"
	"golang.org/x/crypto/ssh"
)

// PowerConnect ...
type PowerConnect struct {
	Driver driver.Factory
	Prompt string
	base   universal.Device
}

// Connect ...
func (dev *PowerConnect) Connect(retries int) error {
	return dev.base.Connect(retries)
}

// Disconnect ...
func (dev *PowerConnect) Disconnect() {
	dev.base.Disconnect()
}

// SendCmd ...
func (dev *PowerConnect) SendCmd(cmd string) (string, error) {
	return dev.base.SendCmd(cmd)
}

// NewClientConfig ...
func (dev PowerConnect) NewClientConfig() *ssh.ClientConfig {
	return nil
}

// NewClient ...
func (dev PowerConnect) NewClient() (*ssh.Client, error) {
	return nil, nil
}

// SendConfig ...
func (dev *PowerConnect) SendConfig(cmds []string) (string, error) {
	return dev.base.SendConfig(cmds)
}
