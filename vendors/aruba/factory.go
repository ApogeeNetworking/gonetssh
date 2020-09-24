package aruba

import (
	"github.com/ApogeeNetworking/gonetssh/client"
	"github.com/ApogeeNetworking/gonetssh/driver"
	"github.com/ApogeeNetworking/gonetssh/universal"
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
