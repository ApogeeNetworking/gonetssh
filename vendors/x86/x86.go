package x86

import (
	"github.com/ApogeeNetworking/gonetssh/driver"
	"github.com/ApogeeNetworking/gonetssh/universal"
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
func (d *PizzaBox) NewClient(sshCfg *ssh.ClientConfig) (*ssh.Client, error) {
	return d.base.NewClient(sshCfg)
}

// NewClientConfig ...
func (d *PizzaBox) NewClientConfig() *ssh.ClientConfig {
	return d.base.NewClientConfig()
}

// SendConfig ...
func (d *PizzaBox) SendConfig(cmds []string) (string, error) {
	return d.base.SendConfig(cmds)
}
