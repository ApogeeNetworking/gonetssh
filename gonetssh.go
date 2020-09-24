package gonetssh

import (
	"strings"

	"github.com/ApogeeNetworking/gonetssh/client"
	"github.com/ApogeeNetworking/gonetssh/universal"
	"github.com/ApogeeNetworking/gonetssh/vendors/aruba"
	"github.com/ApogeeNetworking/gonetssh/vendors/cisco"
	"github.com/ApogeeNetworking/gonetssh/vendors/dell"
	"github.com/ApogeeNetworking/gonetssh/vendors/x86"
)

// DeviceType ...
type DeviceType string

type dType struct {
	CiscoIOS         DeviceType
	CiscoIOSXE       DeviceType
	CiscoAireos      DeviceType
	Aruba            DeviceType
	DellPowerConnect DeviceType
	X86              DeviceType
}

// DType represents a driverType DeviceDriver ENUM
var DType = dType{
	CiscoIOS:         "cisco_ios",
	CiscoIOSXE:       "cisco_iosxe",
	CiscoAireos:      "cisco_aireos",
	Aruba:            "arubaos_ssh",
	DellPowerConnect: "dell_powerconnect",
	X86:              "x86",
}

// NewDevice ...
func NewDevice(host, user, pass, enablePass string, deviceType DeviceType) (universal.Device, error) {
	var device universal.Device
	var err error
	conn, _ := client.NewConnection(host, user, pass)
	switch {
	case strings.Contains(string(deviceType), "cisco"):
		device, err = cisco.NewDevice(conn, string(deviceType), enablePass)
	case strings.Contains(string(deviceType), "dell"):
		device, err = dell.NewDevice(conn, string(deviceType), enablePass)
	case strings.Contains(string(deviceType), "aruba"):
		device, err = aruba.NewDevice(conn, string(deviceType), enablePass)
	case strings.Contains(string(deviceType), "x86"):
		device, err = x86.NewDevice(conn, string(deviceType))
	}
	if err != nil {
		return nil, err
	}
	return device, nil
}
