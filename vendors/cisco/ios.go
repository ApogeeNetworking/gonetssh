package cisco

import (
	"github.com/drkchiloll/gonetssh/driver"
	"github.com/drkchiloll/gonetssh/universal"
	"golang.org/x/crypto/ssh"
)

// IOS ...
type IOS struct {
	Driver driver.Factory
	Prompt string
	base   universal.Device
}

// Connect ...
func (dev *IOS) Connect(retries int) error {
	return dev.base.Connect(retries)
}

// Disconnect ...
func (dev *IOS) Disconnect() {
	dev.base.Disconnect()
}

// SendCmd ...
func (dev *IOS) SendCmd(cmd string) (string, error) {
	return dev.base.SendCmd(cmd)
}

// NewClientConfig ...
func (dev IOS) NewClientConfig() *ssh.ClientConfig {
	return nil
}

// NewClient ...
func (dev IOS) NewClient() (*ssh.Client, error) {
	return nil, nil
}

// SendConfig ...
func (dev *IOS) SendConfig(cmd string) (string, error) {
	return dev.base.SendConfig(cmd)
}
