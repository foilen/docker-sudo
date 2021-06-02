package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

func showPs(containerNames []string) {
	cmd := exec.Command("/usr/bin/docker", "ps", "-a")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	isHeader := true
	scanner := bufio.NewScanner(stdout)
	containerNameRegexp, err := regexp.Compile("([^\\s]+)$")
	for scanner.Scan() {
		if isHeader {
			fmt.Println(scanner.Text())
			isHeader = false
		} else {
			line := scanner.Text()
			containerName := containerNameRegexp.FindString(line)
			if arrayContains(containerNames, containerName) {
				fmt.Println(line)
			}
		}
	}
}
