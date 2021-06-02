package main

import (
	"log"
	"syscall"
)

func showLogs(containerName string) {
	err := syscall.Exec("/usr/bin/docker", []string{"docker", "logs", containerName}, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func showLogsTail(containerName string) {
	err := syscall.Exec("/usr/bin/docker", []string{"docker", "logs", "-f", containerName}, nil)
	if err != nil {
		log.Fatal(err)
	}
}
