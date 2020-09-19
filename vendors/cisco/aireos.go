package cisco

import (
	"github.com/drkchiloll/gonetmiko/driver"
	"github.com/drkchiloll/gonetmiko/types"
	"golang.org/x/crypto/ssh"
)

// AireOS ...
type AireOS struct {
	Driver driver.Factory
	Prompt string
	base   types.Device
}

// Connect ...
func (d *AireOS) Connect(retries int) error {
	return d.base.Connect(retries)
}

// Disconnect ...
func (d *AireOS) Disconnect() {
	d.base.Disconnect()
}

// SendCmd ...
func (d *AireOS) SendCmd(cmd string) (string, error) {
	return d.base.SendCmd(cmd)
}

// NewClientConfig ...
func (d *AireOS) NewClientConfig() *ssh.ClientConfig {
	return nil
}

// NewClient ...
func (d *AireOS) NewClient() (*ssh.Client, error) {
	return nil, nil
}
