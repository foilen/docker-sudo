package main

import (
	"log"
	"syscall"
)

func showRestart(containerName string) {
	err := syscall.Exec("/usr/bin/docker", []string{"docker", "restart", containerName}, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func showStart(containerName string) {
	err := syscall.Exec("/usr/bin/docker", []string{"docker", "start", containerName}, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func showStop(containerName string) {
	err := syscall.Exec("/usr/bin/docker", []string{"docker", "stop", containerName}, nil)
	if err != nil {
		log.Fatal(err)
	}
}
