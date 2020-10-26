# Gonetssh

Gonetssh is a `Golang` based multi-vendor library to simplify SSH Connections to Network(ed) Devices

## Tested On

* Cisco IOS
* Cisco IOS XE
* Cisco NX-OS
* Cisco AireOS WLC
* ArubaOS 6.x WLC
* HP Procurve
* Dell OS6
* Dell PowerConnect
* X86 (SFTP Ops)

Module uses an **"enumerator"** type for supported devices

```go
// DType represents a driverType DeviceDriver ENUM
var DType = dType{
	CiscoIOS:         "cisco_ios",
	CiscoIOSXE:       "cisco_iosxe",
	Cisco9800:        "cisco_9800",
	CiscoAireos:      "cisco_aireos",
	Aruba:            "arubaos_ssh",
	HPProcurve:       "hp_procurve",
	DellOS6:          "dell_os6",
	DellPowerConnect: "dell_pc",
	X86:              "x86",
}
```

## Install

`go get github.com/ApogeeNetworking/gonetssh`

## Usage

**Example 1**: Connect to SFTP Server and Transfer a File

```go
device, _ := gonetssh.NewDevice(
    "host_ip",
    "host_user",
    "host_pass",
    "blank_enable_password",
    gonetssh.DType.X86,
)
// Retrieve SSH Client Configuration so we can EST SSH Client
sshConfig := device.NewClientConfig()
// This EST an SSH Connection for US
sshClient, err := device.NewClient(sshConfig)
if err != nil {
    log.Fatalf("error est ssh connection: %v", err)
}
// Cleanup the Connection when we are done with EXEC
defer device.Disconnect()
scp, err := sftp.NewClient(sshClient)
if err != nil {
    log.Fatalf("error est sftp client conn: %v", err)
}
// Cleanup SCP Connection when we are done with EXEC
defer scp.Close()
// Disregard ERR's for example
destFile, _ := scp.Create("file_to_create")
// Close Destination File when done Writing
defer destFile.Close()
// Disregard error here
sourceFile, _ := os.Open("src_file_to_transfer")
// Close the Source File when Transfer is completed
defer sourceFile.Close()
// Transfer File to Destination
bytesTransferred, _ := io.Copy(destFile, sourceFile)
// Bytes Transferred is an INT64
```