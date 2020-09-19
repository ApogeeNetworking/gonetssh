package cisco

import (
	"fmt"
	"time"

	"github.com/drkchiloll/gonetmiko/client"
	"github.com/drkchiloll/gonetmiko/driver"
	"github.com/drkchiloll/gonetmiko/types"
)

// NewDevice ...
func NewDevice(client client.Connectioner, deviceType, enablePass string) (types.Device, error) {
	prompt := "[[:alnum:]]>.?$|[[:alnum:]]#.?$|[[:alnum:]]\\$.?$"
	driver := driver.NewDriver(client)
	base := BaseDevice{
		Driver:     driver,
		Prompt:     prompt,
		DeviceType: deviceType,
		Delay:      500 * time.Millisecond,
		EnablePass: enablePass,
	}
	switch deviceType {
	case "cisco_ios":
		return &IOS{
			Driver: driver,
			Prompt: prompt,
			base:   &base,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported device type: %s", deviceType)
	}
}
