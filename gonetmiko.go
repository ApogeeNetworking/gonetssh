package gonetmiko

import (
	"strings"

	"github.com/drkchiloll/gonetmiko/client"
	"github.com/drkchiloll/gonetmiko/universal"
	"github.com/drkchiloll/gonetmiko/vendors/aruba"
	"github.com/drkchiloll/gonetmiko/vendors/cisco"
	"github.com/drkchiloll/gonetmiko/vendors/x86"
)

// NewDevice ...
func NewDevice(host, user, pass, enablePass string, deviceType DeviceType) (universal.Device, error) {
	var device universal.Device
	client, err := client.NewConnection(host, user, pass)
	if err != nil {
	}
	switch {
	case strings.Contains(string(deviceType), "cisco"):
		device, err = cisco.NewDevice(client, string(deviceType), enablePass)
	case strings.Contains(string(deviceType), "aruba"):
		device, err = aruba.NewDevice(client, string(deviceType), enablePass)
	case strings.Contains(string(deviceType), "x86"):
		device, err = x86.NewDevice(client, string(deviceType))
	}
	return device, nil
}

// DeviceType ...
type DeviceType string

type dType struct {
	CiscoIOS    DeviceType
	CiscoIOSXE  DeviceType
	CiscoAireos DeviceType
	Aruba       DeviceType
	Dell        DeviceType
	X86         DeviceType
}

// DType represents a driverType DeviceDriver ENUM
var DType = dType{
	CiscoIOS:    "cisco_ios",
	CiscoIOSXE:  "cisco_iosxe",
	CiscoAireos: "cisco_aireos",
	Aruba:       "arubaos_ssh",
	Dell:        "dell",
	X86:         "x86",
}
