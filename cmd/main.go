package main

import (
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
	device, err := gonetssh.NewDevice(
		sshHost,
		sshUser,
		sshPass,
		enablePass,
		gonetssh.DType.DellOS6,
	)
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = device.Connect(20)
	if err != nil {
		log.Fatal(err)
	}
	defer device.Disconnect()
}
