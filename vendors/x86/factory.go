package x86

import (
	"github.com/drkchiloll/gonetmiko/conn"
	"github.com/drkchiloll/gonetmiko/driver"
	"github.com/drkchiloll/gonetmiko/types"
)

// NewDevice instantiates a new X86 for SFTP Mainly
func NewDevice(conn conn.Connectioner, deviceType string) (types.X86, error) {
	driver := driver.NewDriver(conn)
	base := BaseDevice{
		Driver:       driver,
		DeviceType:   deviceType,
		SSHClientCfg: driver.NewClientConfig(),
	}
	return &PizzaBox{
		Driver: driver,
		base:   &base,
	}, nil
}
