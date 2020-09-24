package dell

import (
	"github.com/ApogeeNetworking/gonetssh/driver"
	"github.com/ApogeeNetworking/gonetssh/universal"
	"golang.org/x/crypto/ssh"
)

// OS6 ...
type OS6 struct {
	Driver driver.Factory
	Prompt string
	base   universal.Device
}

// Connect ...
func (dev *OS6) Connect(retries int) error {
	return dev.base.Connect(retries)
}

// Disconnect ...
func (dev *OS6) Disconnect() {
	dev.base.Disconnect()
}

// SendCmd ...
func (dev *OS6) SendCmd(cmd string) (string, error) {
	return dev.base.SendCmd(cmd)
}

// NewClientConfig ...
func (dev OS6) NewClientConfig() *ssh.ClientConfig {
	return nil
}

// NewClient ...
func (dev OS6) NewClient() (*ssh.Client, error) {
	return nil, nil
}

// SendConfig ...
func (dev *OS6) SendConfig(cmds []string) (string, error) {
	return dev.base.SendConfig(cmds)
}
