package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ApogeeNetworking/gonetssh"
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
		gonetssh.DType.Cisco9800,
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
	out, _ := dev.SendCmd("show run")
	fmt.Println(out)
	fmt.Println("running it again")
	out, _ = dev.SendCmd("show run")
	fmt.Println(out)
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
	cmds := []string{
		"save config",
	}
	dev.SendConfig(cmds)
}

func sftpUploadFileExample() {
	dev, _ := gonetssh.NewDevice(
		sshHost,
		sshUser,
		sshPass,
		enablePass,
		gonetssh.DType.X86,
	)
	sshCfg := dev.NewClientConfig()
	sshClient, err := dev.NewClient(sshCfg)
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
