package driver

import (
	"fmt"
	"regexp"
	"time"

	"github.com/ApogeeNetworking/gonetssh/client"
	"golang.org/x/crypto/ssh"
)

// Driver ...
type Driver struct {
	Connection client.Connectioner
}

// Connect ...
func (d *Driver) Connect(retries int) error {
	return d.Connection.Connect(retries)
}

// Disconnect ...
func (d *Driver) Disconnect() {
	d.Connection.Disconnect()
}

// SendCmd ...
func (d *Driver) SendCmd(cmd, prompt string, delay time.Duration) (string, error) {
	d.Connection.Write(cmd + "\n")
	// time.Sleep(delay)
	return d.ReadUntil(prompt)
}

// ReadUntil ...
func (d *Driver) ReadUntil(regex string) (string, error) {
	var result string
	input := make(chan *string)
	stop := make(chan struct{})
	go d.read(regex, input, stop)
	for {
		select {
		case output := <-input:
			switch {
			case output == nil:
				continue
			case output != nil:
				result = *output
			default:
				result = *output
			}
			return result, nil
		case <-stop:
			fmt.Println(<-stop)
			d.Disconnect()
			return "", fmt.Errorf("EOF: %v", <-stop)
		case <-time.After(60 * time.Second):
			d.Connection.Disconnect()
			return "", fmt.Errorf("timeout")
		}
	}
}

func (d *Driver) read(prompt string, buf chan *string, stop chan struct{}) {
	re := regexp.MustCompile(prompt)
	var input string
	for {
		output, err := d.Connection.Read()
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("ERROR", err)
			}
			stop <- struct{}{}
		}
		input += output
		if (len(input) >= 50 && re.MatchString(input[len(input)-45:])) ||
			(len(input) < 50 && re.MatchString(input)) {
			break
		}
		// KeepAlive
		buf <- nil
	}
	buf <- &input
}

// NewClientConfig ...
func (d *Driver) NewClientConfig() *ssh.ClientConfig {
	return d.Connection.NewClientConfig()
}

// NewClient ...
func (d *Driver) NewClient(cfg *ssh.ClientConfig) (*ssh.Client, error) {
	return d.Connection.NewClient(cfg)
}
