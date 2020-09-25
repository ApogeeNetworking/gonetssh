package cisco

import (
	"fmt"
	"regexp"
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
func (d *BaseDevice) SendConfig(cmds []string) (string, error) {
	// Prompt used for any Configure MODE
	prompt := "[[:alnum:]]\\(\\S+\\)#"
	// Pre-pend Relevant Configure Terminal and Exit Commands
	preCmds := []string{"configure terminal"}
	// End the Configuration MODE (CMD)
	cmds = append(cmds, "end")
	preCmds = append(preCmds, cmds...)
	var output string
	for _, cmd := range preCmds {
		switch cmd {
		case "end":
			prompt = d.prompt
		}
		out, _ := d.Driver.SendCmd(cmd, prompt, d.delay)
		output += out
		time.Sleep(d.delay)
	}
	return output, nil
}

// iosPrep ...
func (d *BaseDevice) iosPrep() error {
	// On Cisco_IOS and Cisco_IOSXE set the terminal length for the session
	out, _ := d.SendCmd("terminal len 0")
	// Check whether or not the Prompt is in Exec Mode #
	switch {
	// For Cisco Cat 9800 we want in EXEC mode to push configs (potentially)
	case d.DeviceType == "cisco_9800":
		re := regexp.MustCompile(`[[:alnum:]]>.?$`)
		if re.MatchString(out) {
			d.Driver.ExecEnable(d.EnablePass)
		}
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
