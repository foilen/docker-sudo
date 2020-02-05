package main

import (
	"log"
	"syscall"
)

func showBash(containerName string, userId string) {
	err := syscall.Exec("/usr/bin/docker", []string{"docker", "exec", "-ti", "-u", userId, containerName, "/bin/bash"}, nil)
	if err != nil {
		log.Fatal(err)
	}
}
