package dell

import (
	"time"

	"github.com/drkchiloll/gonetmiko/driver"
	"golang.org/x/crypto/ssh"
)

// BaseDevice represent a generic cisco network object
type BaseDevice struct {
	Driver     driver.Factory
	DeviceType string
	prompt     string
	EnablePass string
	delay      time.Duration
}

// Connect ...
func (d *BaseDevice) Connect(retries int) error {
	if err := d.Driver.Connect(retries); err != nil {
		return err
	}
	switch d.DeviceType {
	case "dell_powerconnect":
		d.prompt = "[[:alnum:]]>.?$|[[:alnum:]]#.?$|[[:alnum:]]\\$.?$"
		d.delay = 2000 * time.Millisecond
		return d.pcPrep()
	default:
		return nil
	}
}

// Disconnect ...
func (d *BaseDevice) Disconnect() {
	d.Driver.Disconnect()
}

// SendCmd ...
func (d *BaseDevice) SendCmd(cmd string) (string, error) {
	return d.Driver.SendCmd(cmd, d.prompt, d.delay)
}

// iosPrep ...
func (d *BaseDevice) pcPrep() error {
	// Set the terminal length for the session
	d.SendCmd("terminal len 0")
	return nil
}

// NewClient ...
func (d *BaseDevice) NewClient() (*ssh.Client, error) {
	return nil, nil
}

// NewClientConfig ...
func (d *BaseDevice) NewClientConfig() *ssh.ClientConfig {
	return nil
}
