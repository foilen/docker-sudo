package main

import (
	"log"
	"syscall"
)

func showBash(containerName string, userId string) {

	tty := "-i"
	if isTTY() {
		tty = "-ti"
	}

	err := syscall.Exec("/usr/bin/docker", []string{"docker", "exec", tty, "-u", userId, containerName, "/bin/bash"}, nil)
	if err != nil {
		log.Fatal(err)
	}
}
