package dell

import (
	"fmt"

	"github.com/drkchiloll/gonetmiko/client"
	"github.com/drkchiloll/gonetmiko/driver"
	"github.com/drkchiloll/gonetmiko/universal"
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
