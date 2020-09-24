package dell

import (
	"fmt"

	"github.com/ApogeeNetworking/gonetssh/client"
	"github.com/ApogeeNetworking/gonetssh/driver"
	"github.com/ApogeeNetworking/gonetssh/universal"
)

// NewDevice ...
func NewDevice(conn client.Connectioner, deviceType, enablePass string) (universal.Device, error) {
	driver := driver.NewDriver(conn)
	base := BaseDevice{
		Driver:     driver,
		DeviceType: deviceType,
		EnablePass: enablePass,
	}
	switch deviceType {
	case "dell_powerconnect":
		return &PowerConnect{
			Driver: driver,
			base:   &base,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported device type: %s", deviceType)
	}
}
