package x86

import (
	"github.com/drkchiloll/gonetmiko/driver"
	"github.com/drkchiloll/gonetmiko/universal"
	"golang.org/x/crypto/ssh"
)

// PizzaBox ...
type PizzaBox struct {
	Driver       driver.Factory
	base         universal.Device
	SSHClientCfg *ssh.ClientConfig
}

// Connect ...
func (d *PizzaBox) Connect(retries int) error {
	return d.base.Connect(retries)
}

// Disconnect ...
func (d *PizzaBox) Disconnect() {
	d.base.Disconnect()
}

// SendCmd ...
func (d *PizzaBox) SendCmd(cmd string) (string, error) {
	return "", nil
}

// NewClient ...
func (d *PizzaBox) NewClient() (*ssh.Client, error) {
	return d.base.NewClient()
}

// NewClientConfig ...
func (d *PizzaBox) NewClientConfig() *ssh.ClientConfig {
	return d.base.NewClientConfig()
}
