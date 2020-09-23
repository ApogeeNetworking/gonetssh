package cisco

import (
	"fmt"
	"regexp"
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
	switch d.DeviceType {
	case "cisco_ios":
		d.prompt = "[[:alnum:]]>.?$|[[:alnum:]]#.?$|[[:alnum:]]\\$.?$"
		d.delay = 250 * time.Millisecond
		return d.iosPrep()
	case "cisco_aireos":
		d.prompt = `\s>.?$`
		d.delay = 500 * time.Millisecond
		return d.aireosPrep()
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

// SendConfig ...
func (d *BaseDevice) SendConfig(cmd string) (string, error) {
	return d.Driver.SendCmd(cmd, d.prompt, d.delay)
}

// iosPrep ...
func (d *BaseDevice) iosPrep() error {
	// On Cisco_IOS and Cisco_IOSXE set the terminal length for the session
	out, _ := d.SendCmd("terminal len 0")
	// Check whether or not the Prompt is in Exec Mode #
	re := regexp.MustCompile(`[[:alnum:]]>.?$`)
	if re.MatchString(out) {
		fmt.Println("wasn't in enable mode")
		d.Driver.ExecEnable(d.EnablePass)
	}
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
