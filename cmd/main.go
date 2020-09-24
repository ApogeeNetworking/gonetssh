package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/drkchiloll/gonetssh"
	"github.com/pkg/sftp"
	"github.com/subosito/gotenv"
)

var sshHost, sshUser, sshPass, enablePass string

func init() {
	gotenv.Load()
	sshHost = os.Getenv("SSH_HOST")
	sshUser = os.Getenv("SSH_USER")
	sshPass = os.Getenv("SSH_PASS")
	enablePass = os.Getenv("ENABLE_PW")
}

func main() {
	dev, err := gonetssh.NewDevice(
		sshHost,
		sshUser,
		sshPass,
		enablePass,
		gonetssh.DType.CiscoIOS,
	)
	if err != nil {
		// Device Type Not Supported
		log.Fatalf("%v", err)
	}
	err = dev.Connect(10)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer dev.Disconnect()
	orgCmds := []string{
		"ap 6c41.0ec7.cafc",
		"policy-tag APG_Better-AP-Group",
		"site-tag Better-AP-Group",
		"rf-tag \"mga - Default\"",
	}
	// orgCmds := []string{
	// 	"ap 6c41.0ec7.cafc",
	// 	"policy-tag \"No SSID\"",
	// 	"rf-tag \"No SSID\"",
	// 	"site-tag default-site-tag",
	// }
	dev.SendConfig(orgCmds)
}

func aireOS() {
	dev, _ := gonetssh.NewDevice(
		sshHost,
		sshUser,
		sshPass,
		enablePass,
		gonetssh.DType.CiscoAireos,
	)
	err := dev.Connect(10)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer dev.Disconnect()
	res, err := dev.SendCmd("show ap inventory all")
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(res)
}

func sftpUploadFileExample() {
	dev, _ := gonetssh.NewDevice(
		sshHost,
		sshUser,
		sshPass,
		enablePass,
		gonetssh.DType.X86,
	)
	dev.NewClientConfig()
	sshClient, err := dev.NewClient()
	if err != nil {
		log.Fatalf("%v", err)
	}
	scp, err := sftp.NewClient(sshClient)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer func() {
		scp.Close()
		dev.Disconnect()
	}()
	dstFile, err := scp.Create("gonetssh.go")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer dstFile.Close()
	srcFile, err := os.Open("gonetssh.go")
	if err != nil {
		log.Fatalf("%v", err)
	}
	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("%d bytes copied\n", bytes)
}

func ios() {
	dev, _ := gonetssh.NewDevice(
		sshHost,
		sshUser,
		sshPass,
		enablePass,
		gonetssh.DType.CiscoIOS,
	)
	err := dev.Connect(10)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer dev.Disconnect()
	res, err := dev.SendCmd("show cdp neighbor")
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(res)
}
