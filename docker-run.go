package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func showRun(imageName string, userId string, userHome string) {

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
	err = syscall.Exec("/usr/bin/docker", []string{"docker", "run", "-ti", "--rm", "-u", userId, "-v", userHome + ":" + userHome, uuid, "/bin/bash"}, nil)
	if err != nil {
		log.Fatal(err)
	}

}
