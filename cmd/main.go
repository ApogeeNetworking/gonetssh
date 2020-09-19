package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/drkchiloll/gonetmiko"
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

func main() {}

func sftpUploadFileExample() {
	dev, _ := gonetmiko.NewDevice(
		sshHost,
		sshUser,
		sshPass,
		enablePass,
		gonetmiko.DType.X86,
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
	dstFile, err := scp.Create("gonetmiko.go")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer dstFile.Close()
	srcFile, err := os.Open("gonetmiko.go")
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
	dev, _ := gonetmiko.NewDevice(
		sshHost,
		sshUser,
		sshPass,
		enablePass,
		gonetmiko.DType.CiscoIOS,
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
