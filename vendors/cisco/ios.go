package cisco

import (
	"regexp"
	"strings"

	"github.com/ApogeeNetworking/gonetssh/driver"
	"github.com/ApogeeNetworking/gonetssh/universal"
	"golang.org/x/crypto/ssh"
)

// IOS ...
type IOS struct {
	Driver     driver.Factory
	base       universal.Device
	prompt     string
	deviceType string
}

// Connect ...
func (dev *IOS) Connect(retries int) error {
	return dev.base.Connect(retries)
}

// Disconnect ...
func (dev *IOS) Disconnect() {
	dev.base.Disconnect()
}

// SendCmd ...
func (dev *IOS) SendCmd(cmd string) (string, error) {
	out, err := dev.base.SendCmd(cmd)
	switch {
	case dev.deviceType == "cisco_ios":
		return dev.processIosSendCmd(cmd, out)
	default:
		return out, err
	}
}

func (dev *IOS) processIosSendCmd(cmd, out string) (string, error) {
	lines := strings.Split(out, "\n")
	var output string
	prRe := regexp.MustCompile(dev.prompt)
	for _, line := range lines {
		if prRe.MatchString(line) || line == "" ||
			strings.Contains(line, cmd) {
			continue
		}
		output += line + "\n"
	}
	return output, nil
}

// NewClientConfig ...
func (dev IOS) NewClientConfig() *ssh.ClientConfig {
	return nil
}

// NewClient ...
func (dev IOS) NewClient(sshCfg *ssh.ClientConfig) (*ssh.Client, error) {
	return nil, nil
}

// SendConfig ...
func (dev *IOS) SendConfig(cmds []string) (string, error) {
	return dev.base.SendConfig(cmds)
}
