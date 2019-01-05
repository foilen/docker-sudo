package main

import (
	"log"
	"syscall"
)

func showStop(containerName string) {
	err := syscall.Exec("/usr/bin/docker", []string{"docker", "stop", containerName}, nil)
	if err != nil {
		log.Fatal(err)
	}
}
