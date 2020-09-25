package cisco

import (
	"fmt"

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
	switch {
	case deviceType == "cisco_ios" || deviceType == "cisco_9800":
		return &IOS{
			Driver: driver,
			base:   &base,
		}, nil
	case deviceType == "cisco_aireos":
		return &AireOS{
			Driver: driver,
			base:   &base,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported device type: %s", deviceType)
	}
}
