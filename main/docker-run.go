package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func showRunCustomImage(customImageName string, userId string, userHome string, network string) {

	// Get a uuid to use as a tag
	uuid, err := getCommandOutput("/usr/bin/uuidgen")
	if err != nil {
		log.Fatalln(err)
	}
	uuid = strings.Trim(uuid, "\n")

	// Build
	cmd := exec.Command("/usr/bin/docker", "build", "-t", uuid, ".")
	cmd.Dir = "/etc/docker-sudo/images/" + customImageName
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
	err = syscall.Exec("/usr/bin/docker", []string{"docker", "run", tty, "--rm", "-u", userId, "-v", userHome + ":" + userHome, "-w", workDir, "--network", network, uuid, "/bin/bash"}, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func showRunImage(imageName string, userId string, userHome string, network string) {

	// Run
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	tty := "-i"
	if isTTY() {
		tty = "-ti"
	}
	err = syscall.Exec("/usr/bin/docker", []string{"docker", "run", tty, "--rm", "-u", userId, "-v", userHome + ":" + userHome, "-w", workDir, "--network", network, imageName, "/bin/bash"}, nil)
	if err != nil {
		log.Fatal(err)
	}

}
