package aruba

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
	switch d.DeviceType {
	case "aruba6_ssh":
		d.prompt = `\)\s#.?$`
	case "aruba8_ssh":
		d.prompt = `\)\s\*#.?$`
	}
	d.delay = 1000 * time.Millisecond
	return d.arubaPrep()
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
	}
	return output, nil
}

// arubaPrep in order to do anything useful on the Aruba CLI
// you must enter EXEC Mode by issuing the enable command
// once in EXEC mode the no paging command is entered in order
// to get all data from command on issuance (MINUS --More-- keyword)
func (d *BaseDevice) arubaPrep() error {
	if d.DeviceType == "aruba6_ssh" {
		// Temporary Enable Password Prompt
		prompt := "Password:"
		// Enter Enable Command; the Prompt will be Password:
		d.Driver.SendCmd("enable", prompt, d.delay)
		// Enter Enable Password as CMD (as that's what the input really is)
		// Change the Prompt Back to (controller) # = d.prompt
		d.Driver.SendCmd(d.EnablePass, d.prompt, d.delay)
	}
	// set the terminal length for the session
	d.SendCmd("no paging")
	return nil
}

// NewClient ...
func (d *BaseDevice) NewClient(sshCfg *ssh.ClientConfig) (*ssh.Client, error) {
	return nil, nil
}

// NewClientConfig ...
func (d *BaseDevice) NewClientConfig() *ssh.ClientConfig {
	return nil
}
