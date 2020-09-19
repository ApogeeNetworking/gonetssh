package gonetmiko

import (
	"strings"

	"github.com/drkchiloll/gonetmiko/conn"
	"github.com/drkchiloll/gonetmiko/types"
	"github.com/drkchiloll/gonetmiko/vendors/cisco"
)

// NewDevice ...
func NewDevice(host, user, pass, enablePass string, deviceType DeviceType) (types.Device, error) {
	var device types.Device
	conn, err := conn.NewConnection(host, user, pass)
	if err != nil {
	}
	switch {
	case strings.Contains(string(deviceType), "cisco"):
		device, err = cisco.NewDevice(conn, string(deviceType), enablePass)
	}
	return device, nil
}

// DeviceType ...
type DeviceType string

type dType struct {
	CiscoIOS    DeviceType
	CiscoIOSXE  DeviceType
	Aruba       DeviceType
	CiscoAireos DeviceType
	Dell        DeviceType
	X86         DeviceType
}

// DType represents a driverType DeviceDriver ENUM
var DType = dType{
	CiscoIOS:    "cisco_ios",
	CiscoIOSXE:  "cisco_iosxe",
	Aruba:       "aruba",
	CiscoAireos: "cisco_wlc",
	Dell:        "dell",
	X86:         "x86",
}
