package cisco

import (
	"fmt"

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
	switch deviceType {
	case "cisco_ios":
		return &IOS{
			Driver: driver,
			base:   &base,
		}, nil
	case "cisco_aireos":
		return &AireOS{
			Driver: driver,
			base:   &base,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported device type: %s", deviceType)
	}
}
