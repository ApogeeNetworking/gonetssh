package hp

import (
	"github.com/ApogeeNetworking/gonetssh/driver"
	"github.com/ApogeeNetworking/gonetssh/universal"
	"golang.org/x/crypto/ssh"
)

// SSH ...
type SSH struct {
	Driver driver.Factory
	Prompt string
	base   universal.Device
}

// Connect ...
func (dev *SSH) Connect(retries int) error {
	return dev.base.Connect(retries)
}

// Disconnect ...
func (dev *SSH) Disconnect() {
	dev.base.Disconnect()
}

// SendCmd ...
func (dev *SSH) SendCmd(cmd string) (string, error) {
	return dev.base.SendCmd(cmd)
}

// NewClientConfig ...
func (dev SSH) NewClientConfig() *ssh.ClientConfig {
	return nil
}

// NewClient ...
func (dev SSH) NewClient(sshCfg *ssh.ClientConfig) (*ssh.Client, error) {
	return nil, nil
}

// SendConfig ...
func (dev *SSH) SendConfig(cmds []string) (string, error) {
	return dev.base.SendConfig(cmds)
}
