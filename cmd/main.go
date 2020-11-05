package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/ApogeeNetworking/gonetssh"
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

func trimWS(text string) string {
	tsRe := regexp.MustCompile(`\s+`)
	return tsRe.ReplaceAllString(text, " ")
}

func main() {
	dev, err := gonetssh.NewDevice(
		sshHost,
		sshUser,
		sshPass,
		enablePass,
		gonetssh.DType.Aruba8,
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

	out, _ := dev.SendCmd("show configuration effective | include essid")
	fmt.Println(out)
}
