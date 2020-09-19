package main

import (
	"fmt"
	"log"
	"os"

	"github.com/drkchiloll/gonetmiko"
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
