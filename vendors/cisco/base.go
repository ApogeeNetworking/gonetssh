package cisco

import (
	"regexp"
	"time"

	"github.com/drkchiloll/gonetmiko/driver"
	"golang.org/x/crypto/ssh"
)

// BaseDevice represent a generic cisco network object
type BaseDevice struct {
	Driver     driver.Factory
	DeviceType string
	Prompt     string
	EnablePass string
	Delay      time.Duration
}

// Connect ...
func (d *BaseDevice) Connect(retries int) error {
	if err := d.Driver.Connect(retries); err != nil {
		return err
	}
	d.Prompt = "[[:alnum:]]>.?$|[[:alnum:]]#.?$|[[:alnum:]]\\$.?$"
	return d.sessionPrep()
}

// Disconnect ...
func (d *BaseDevice) Disconnect() {
	d.Driver.Disconnect()
}

// SendCmd ...
func (d *BaseDevice) SendCmd(cmd string) (string, error) {
	return d.Driver.SendCmd(cmd, d.Prompt, d.Delay)
}

func (d *BaseDevice) sessionPrep() error {
	// Check whether or not the Prompt is in Exec Mode #
	re := regexp.MustCompile(`[[:alnum:]]>.?$`)
	out, _ := d.Driver.SendCmd("terminal len 0", d.Prompt, d.Delay)
	if re.MatchString(out) {
		d.Driver.ExecEnable(d.EnablePass)
	}
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
