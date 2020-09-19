package aruba

import (
	"github.com/drkchiloll/gonetmiko/client"
	"github.com/drkchiloll/gonetmiko/driver"
	"github.com/drkchiloll/gonetmiko/universal"
)

// NewDevice ...
func NewDevice(client client.Connectioner, deviceType, enablePass string) (universal.Device, error) {
	driver := driver.NewDriver(client)
	base := BaseDevice{
		Driver:     driver,
		DeviceType: deviceType,
		EnablePass: enablePass,
	}
	return &SSH{
		Driver: driver,
		base:   &base,
	}, nil
}
