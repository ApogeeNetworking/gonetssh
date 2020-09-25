package dell

import (
	"fmt"

	"github.com/ApogeeNetworking/gonetssh/client"
	"github.com/ApogeeNetworking/gonetssh/driver"
	"github.com/ApogeeNetworking/gonetssh/universal"
)

// NewDevice ...
func NewDevice(conn client.Connectioner, user, pass, deviceType, enablePass string) (universal.Device, error) {
	driver := driver.NewDriver(conn)
	base := BaseDevice{
		Driver:     driver,
		DeviceType: deviceType,
		enablePass: enablePass,
		user:       user,
		pass:       pass,
	}
	switch {
	case deviceType == "dell_os6" || deviceType == "dell_pc":
		return &OS6{
			Driver: driver,
			base:   &base,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported device type: %s", deviceType)
	}
}
