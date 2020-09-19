package x86

import (
	"github.com/drkchiloll/gonetmiko/driver"
	"golang.org/x/crypto/ssh"
)

// BaseDevice represents a generic X86 Server
type BaseDevice struct {
	Driver       driver.Factory
	DeviceType   string
	SSHClientCfg *ssh.ClientConfig
}

// Connect ...
func (d *BaseDevice) Connect(retries int) error {
	if err := d.Driver.Connect(retries); err != nil {
		return err
	}
	return nil
}

// Disconnect ...
func (d *BaseDevice) Disconnect() {
	d.Driver.Disconnect()
}

// SendCmd ...
func (d *BaseDevice) SendCmd(cmd string) (string, error) {
	return "", nil
}

// NewClientConfig ...
func (d *BaseDevice) NewClientConfig() *ssh.ClientConfig {
	return d.Driver.NewClientConfig()
}

// NewClient ...
func (d *BaseDevice) NewClient() (*ssh.Client, error) {
	return d.Driver.NewClient(d.SSHClientCfg)
}