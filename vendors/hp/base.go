package hp

import (
	"time"

	"github.com/ApogeeNetworking/gonetssh/driver"
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
	d.delay = 100 * time.Millisecond
	switch {
	case d.DeviceType == "hp_procurve":
		d.prompt = "[[:alnum:]]#"
		d.Driver.SendCmd("\n", d.prompt, d.delay)
		d.Driver.SendCmd("no page", d.prompt, d.delay)
		return nil
	}
	return nil
}

// Disconnect ...
func (d *BaseDevice) Disconnect() {
	d.Driver.Disconnect()
}

// SendCmd ...
func (d *BaseDevice) SendCmd(cmd string) (string, error) {
	return d.Driver.SendCmd(cmd, d.prompt, d.delay)
}

// SendConfig ...
func (d *BaseDevice) SendConfig(cmds []string) (string, error) {
	// Prompt used for Configure MODE
	var output string
	for _, cmd := range cmds {
		out, _ := d.Driver.SendCmd(cmd, d.prompt, d.delay)
		output += out
		time.Sleep(100 * time.Millisecond)
	}
	return output, nil
}

// NewClient ...
func (d *BaseDevice) NewClient(sshCfg *ssh.ClientConfig) (*ssh.Client, error) {
	return nil, nil
}

// NewClientConfig ...
func (d *BaseDevice) NewClientConfig() *ssh.ClientConfig {
	return nil
}
