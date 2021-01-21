package cisco

import (
	"fmt"

	"github.com/ApogeeNetworking/gonetssh/client"
	"github.com/ApogeeNetworking/gonetssh/driver"
	"github.com/ApogeeNetworking/gonetssh/universal"
)

// NewDevice ...
func NewDevice(client client.Connectioner, user, pass, deviceType, enablePass string) (universal.Device, error) {
	driver := driver.NewDriver(client)
	base := BaseDevice{
		Driver:     driver,
		DeviceType: deviceType,
		User:       user,
		Pass:       pass,
		EnablePass: enablePass,
	}
	switch {
	case deviceType == "cisco_ios" || deviceType == "cisco_9800":
		return &IOS{
			Driver:     driver,
			base:       &base,
			deviceType: deviceType,
			prompt:     "[[:alnum:]]>.?$|[[:alnum:]]#.?$",
		}, nil
	case deviceType == "cisco_aireos" || deviceType == "cisco_aireos_old":
		return &AireOS{
			Driver: driver,
			base:   &base,
			prompt: `\s>.?$`,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported device type: %s", deviceType)
	}
}
