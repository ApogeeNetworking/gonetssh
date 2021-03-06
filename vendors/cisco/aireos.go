package cisco

import (
	"github.com/ApogeeNetworking/gonetssh/driver"
	"github.com/ApogeeNetworking/gonetssh/universal"
	"golang.org/x/crypto/ssh"
)

// AireOS ...
type AireOS struct {
	Driver driver.Factory
	prompt string
	base   universal.Device
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
func (d *AireOS) NewClient(sshCfg *ssh.ClientConfig) (*ssh.Client, error) {
	return nil, nil
}

// SendConfig ...
func (d *AireOS) SendConfig(cmds []string) (string, error) {
	return d.base.SendConfig(cmds)
}
