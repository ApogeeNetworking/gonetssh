package aruba

import (
	"github.com/drkchiloll/gonetssh/client"
	"github.com/drkchiloll/gonetssh/driver"
	"github.com/drkchiloll/gonetssh/universal"
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
