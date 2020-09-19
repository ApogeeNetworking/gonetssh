package x86

import (
	"github.com/drkchiloll/gonetmiko/client"
	"github.com/drkchiloll/gonetmiko/driver"
	"github.com/drkchiloll/gonetmiko/types"
)

// NewDevice instantiates a new X86 for SFTP Mainly
func NewDevice(cl client.Connectioner, deviceType string) (types.Device, error) {
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
