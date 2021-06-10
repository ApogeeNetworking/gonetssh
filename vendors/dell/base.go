package dell

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
	enablePass string
	delay      time.Duration
	user       string
	pass       string
}

// Connect ...
func (d *BaseDevice) Connect(retries int) error {
	if err := d.Driver.Connect(retries); err != nil {
		return err
	}
	d.prompt = "\\S+>.?$|\\S+#.?$"
	d.delay = 250 * time.Millisecond
	switch d.DeviceType {
	case "dell_os6":
		return d.os6Prep()
	case "dell_pc":
		return d.powerConnectPrep()
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
	if cmd == "wr" {
		customPrmpt := `Are\s+you\s+sure\syou\s+want\s+to\s+save\?\s+\(y\/n\)\s+`
		d.Driver.SendCmd(cmd, customPrmpt, d.delay)
		cmd = "y"
	}
	return d.Driver.SendCmd(cmd, d.prompt, d.delay)
}

// os6Prep OS6 Connect Preparation
func (d *BaseDevice) os6Prep() error {
	// Set the terminal length for the session
	d.SendCmd("terminal len 0")
	return nil
}

// pcPrep PowerConnect Prepaparation
func (d *BaseDevice) powerConnectPrep() error {
	// If Connections was a Success Enter User Name (Prompt being Password)
	// out, _ := d.Driver.SendCmd(d.user, `User\s+Name:|Password:`, d.delay)
	d.Driver.SendCmd(d.user, `User\s+Name:`, d.delay)
	// Enter Password with normal Prompt
	d.Driver.SendCmd(d.pass, `.*`, d.delay)
	// // Enter Terminal Length 0 so that it doesn't have to bother with
	d.Driver.SendCmd("terminal datadump", d.prompt, d.delay)
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

// SendConfig ...
func (d *BaseDevice) SendConfig(cmds []string) (string, error) {
	var results string
	for _, cmd := range cmds {
		out, _ := d.SendCmd(cmd)
		time.Sleep(500 * time.Millisecond)
		results += out
	}
	return results, nil
}
