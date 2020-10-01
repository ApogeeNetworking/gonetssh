package cisco

import (
	"regexp"
	"strings"
	"time"

	"github.com/ApogeeNetworking/gonetssh/driver"
	"golang.org/x/crypto/ssh"
)

var contains = strings.Contains

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
	switch {
	case d.DeviceType == "cisco_ios" || d.DeviceType == "cisco_9800":
		d.prompt = "[[:alnum:]]>.?$|[[:alnum:]]#.?$"
		d.delay = 250 * time.Millisecond
		return d.iosPrep()
	case d.DeviceType == "cisco_aireos":
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
	switch {
	case d.DeviceType == "cisco_aireos":
		return d.handleAireosConfigs(cmds[0])
	default:
		// Currently ONLY Catalyst 9800 WLCs
		return d.handleIOSConfigs(cmds)
	}
}

func (d *BaseDevice) handleIOSConfigs(cmds []string) (string, error) {
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

func (d *BaseDevice) handleAireosConfigs(cmd string) (string, error) {
	switch {
	case contains(cmd, "config ap group") ||
		contains(cmd, "clear ap config") ||
		contains(cmd, "save config"):
		// Changing the AP's group name will cause the AP to reboot.
		// Clearing an AP's config (factory reset) causes Reboot of course.
		// Are you sure you want to continue|save? (y/n)
		prompt := `\(y\/n\)`
		d.Driver.SendCmd(cmd, prompt, d.delay)
		return d.Driver.SendCmd("y", d.prompt, d.delay)
	default:
		return d.SendCmd(cmd)
	}
}

// iosPrep ...
func (d *BaseDevice) iosPrep() error {
	// If Device Always goes into EXEC Mode
	// Test By Sending the disable cmd here
	// d.SendCmd("disable")
	// On Cisco_IOS and Cisco_IOSXE set the terminal length for the session
	out, _ := d.SendCmd("terminal len 0")
	// Check whether or not the Prompt is in Exec Mode #
	switch {
	// For Cisco Cat 9800 we want in EXEC mode to push configs (potentially)
	case d.DeviceType == "cisco_9800":
		re := regexp.MustCompile(`[[:alnum:]]>.?$`)
		if re.MatchString(out) {
			prompt := `Password:\s`
			_, _ = d.Driver.SendCmd("enable", prompt, d.delay)
			// For Debugging It Might be useful to make the _ = out
			_, _ = d.Driver.SendCmd(d.EnablePass, d.prompt, d.delay)
			// For Debugging It Might be useful to make the _ = out
		}
	}
	return nil
}

func (d *BaseDevice) aireosPrep() error {
	d.SendCmd("config paging disable")
	return nil
}

// NewClient not implemented for Cisco Devices
func (d *BaseDevice) NewClient() (*ssh.Client, error) {
	return nil, nil
}

// NewClientConfig not implemented for Cisco Devices
func (d *BaseDevice) NewClientConfig() *ssh.ClientConfig {
	return nil
}
