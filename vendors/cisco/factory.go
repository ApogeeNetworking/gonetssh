package cisco

import (
	"fmt"
	"time"

	"github.com/drkchiloll/gonetmiko/conn"
	"github.com/drkchiloll/gonetmiko/driver"
	"github.com/drkchiloll/gonetmiko/types"
)

// NewDevice ...
func NewDevice(conn conn.Connectioner, deviceType, enablePass string) (types.Device, error) {
	prompt := "[[:alnum:]]>.?$|[[:alnum:]]#.?$|[[:alnum:]]\\$.?$"
	driver := driver.NewDriver(conn)
	base := BaseDevice{
		Driver:     driver,
		Prompt:     prompt,
		DeviceType: deviceType,
		Delay:      1 * time.Second,
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
