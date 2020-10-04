package x86

import (
	"github.com/ApogeeNetworking/gonetssh/client"
	"github.com/ApogeeNetworking/gonetssh/driver"
	"github.com/ApogeeNetworking/gonetssh/universal"
)

// NewDevice instantiates a new X86 for SFTP Mainly
func NewDevice(cl client.Connectioner, deviceType string) (universal.Device, error) {
	driver := driver.NewDriver(cl)
	base := BaseDevice{
		Driver:     driver,
		DeviceType: deviceType,
	}
	return &PizzaBox{
		Driver: driver,
		base:   &base,
	}, nil
}
