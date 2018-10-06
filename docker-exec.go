package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func showExec(imageName string, command string, arguments []string, userId string, userHome string) {

	// Get a uuid to use as a tag
	uuid, err := getCommandOutput("/usr/bin/uuidgen")
	if err != nil {
		log.Fatalln(err)
	}
	uuid = strings.Trim(uuid, "\n")

	// Build
	cmd := exec.Command("/usr/bin/docker", "build", "-t", uuid, ".")
	cmd.Dir = "/etc/docker-sudo/images/" + imageName
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Run
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	tty := "-i"
	if isTTY() {
		tty = "-ti"
	}
	argv := []string{"docker", "run", tty, "--rm", "-u", userId, "-v", userHome + ":" + userHome, "-w", workDir, uuid, command}
	if len(arguments) > 0 {
		tmp := make([]string, len(argv)+len(arguments))
		copy(tmp, argv)
		copy(tmp[len(argv):], arguments)
		argv = tmp
	}
	err = syscall.Exec("/usr/bin/docker", argv, nil)
	if err != nil {
		log.Fatal(err)
	}

}
