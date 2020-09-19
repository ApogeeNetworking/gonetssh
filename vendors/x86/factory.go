package x86

import (
	"github.com/drkchiloll/gonetssh/client"
	"github.com/drkchiloll/gonetssh/driver"
	"github.com/drkchiloll/gonetssh/universal"
)

// NewDevice instantiates a new X86 for SFTP Mainly
func NewDevice(cl client.Connectioner, deviceType string) (universal.Device, error) {
	driver := driver.NewDriver(cl)
	base := BaseDevice{
		Driver:       driver,
		DeviceType:   deviceType,
		SSHClientCfg: driver.NewClientConfig(),
	}
	return &PizzaBox{
		Driver: driver,
		base:   &base,
	}, nil
}
