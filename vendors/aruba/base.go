package aruba

import (
	"fmt"
	"time"

	"github.com/drkchiloll/gonetssh/driver"
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
	d.prompt = `\)\s#.?$`
	d.delay = 250 * time.Millisecond
	return d.sessionPrep()
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
	return "", nil
}

// iosPrep ...
func (d *BaseDevice) sessionPrep() error {
	d.Driver.ExecEnable(d.EnablePass)
	// set the terminal length for the session
	d.SendCmd("no paging")
	return nil
}

func (d *BaseDevice) aireosPrep() error {
	out, _ := d.SendCmd("config paging disable")
	fmt.Println(out)

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
