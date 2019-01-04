package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {

	// Check the current user id
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	username := currentUser.Username
	userid := currentUser.Uid
	userHome := currentUser.HomeDir

	// Read the list of available container ids
	containerNames, _ := getContainerList(username)
	customImageNames, _ := getCustomImagesList()

	// Arguments
	if len(os.Args) == 1 {
		showHelp(containerNames, customImageNames)
		os.Exit(1)
	}
	command := os.Args[1]
	if command == "help" {
		showHelp(containerNames, customImageNames)
		os.Exit(0)
	}

	if command == "ps" {
		showPs(containerNames)
		os.Exit(0)
	}

	if command == "logs" {
		if len(os.Args) < 3 {
			showHelp(containerNames, customImageNames)
			os.Exit(1)
		}

		// Check permitted name
		containerName := os.Args[2]
		if arrayContains(containerNames, containerName) {
			showLogs(containerName)
			os.Exit(0)
		} else {
			fmt.Println("You do not have access to container:", containerName)
			showHelp(containerNames, customImageNames)
			os.Exit(1)
		}
	}

	if command == "bash" {
		if len(os.Args) < 3 {
			showHelp(containerNames, customImageNames)
			os.Exit(1)
		}

		// Check permitted name
		containerName := os.Args[2]
		if arrayContains(containerNames, containerName) {
			showBash(containerName, userid)
			os.Exit(0)
		} else {
			fmt.Println("You do not have access to container:", containerName)
			showHelp(containerNames, customImageNames)
			os.Exit(1)
		}
	}

	if command == "tails" {
		if len(os.Args) < 3 {
			showHelp(containerNames, customImageNames)
			os.Exit(1)
		}

		// Check permitted name
		containerName := os.Args[2]
		if arrayContains(containerNames, containerName) {
			showLogsTail(containerName)
			os.Exit(0)
		} else {
			fmt.Println("You do not have access to container:", containerName)
			showHelp(containerNames, customImageNames)
			os.Exit(1)
		}
	}

	if command == "run" {
		if len(os.Args) < 3 {
			showHelp(containerNames, customImageNames)
			os.Exit(1)
		}

		// Check permitted name
		imageName := os.Args[2]
		if arrayContains(customImageNames, imageName) {
			showRun(imageName, userid, userHome)
			os.Exit(0)
		} else {
			fmt.Println("You do not have access to the custom image:", imageName)
			showHelp(containerNames, customImageNames)
			os.Exit(1)
		}
	}

	if command == "exec" {
		if len(os.Args) < 4 {
			showHelp(containerNames, customImageNames)
			os.Exit(1)
		}

		// Check permitted name
		imageName := os.Args[2]
		if arrayContains(customImageNames, imageName) {
			// Get command and arguments
			cmd := os.Args[3]
			arguments := os.Args[4:]
			showExec(imageName, cmd, arguments, userid, userHome)
			os.Exit(0)
		} else {
			fmt.Println("You do not have access to the custom image:", imageName)
			showHelp(containerNames, customImageNames)
			os.Exit(1)
		}
	}

	fmt.Println("Invalid command", command)
	showHelp(containerNames, customImageNames)
	os.Exit(1)
}

func getContainerList(username string) ([]string, error) {
	return getListInFile("/etc/docker-sudo/containers-" + username + ".conf")
}

func getCustomImagesList() ([]string, error) {
	return getListInFolder("/etc/docker-sudo/images")
}
