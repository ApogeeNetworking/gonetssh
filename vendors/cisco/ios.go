package cisco

import (
	"github.com/drkchiloll/gonetmiko/driver"
	"github.com/drkchiloll/gonetmiko/types"
	"golang.org/x/crypto/ssh"
)

// IOS ...
type IOS struct {
	Driver driver.Factory
	Prompt string
	base   types.Device
}

// Connect ...
func (d *IOS) Connect(retries int) error {
	return d.base.Connect(retries)
}

// Disconnect ...
func (d *IOS) Disconnect() {
	d.base.Disconnect()
}

// SendCmd ...
func (d *IOS) SendCmd(cmd string) (string, error) {
	return d.base.SendCmd(cmd)
}

// NewClientConfig ...
func (d *IOS) NewClientConfig() *ssh.ClientConfig {
	return nil
}

// NewClient ...
func (d *IOS) NewClient() (*ssh.Client, error) {
	return nil, nil
}