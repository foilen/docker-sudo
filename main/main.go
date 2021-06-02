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

	// Get the config if available
	var network = "bridge"
	dockerConfiguration, err := getDockerConfiguration("/etc/docker-sudo/config.json")
	if err == nil {
		network = dockerConfiguration.Network
	} else if !os.IsNotExist(err) {
		log.Fatal("Problem loading the file /etc/docker-sudo/config.json: ", err)
	}

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

	if command == "restart" {
		if len(os.Args) < 3 {
			showHelp(containerNames, customImageNames)
			os.Exit(1)
		}

		// Check permitted name
		containerName := os.Args[2]
		if arrayContains(containerNames, containerName) {
			showRestart(containerName)
			os.Exit(0)
		} else {
			fmt.Println("You do not have access to container:", containerName)
			showHelp(containerNames, customImageNames)
			os.Exit(1)
		}
	}

	if command == "start" {
		if len(os.Args) < 3 {
			showHelp(containerNames, customImageNames)
			os.Exit(1)
		}

		// Check permitted name
		containerName := os.Args[2]
		if arrayContains(containerNames, containerName) {
			showStart(containerName)
			os.Exit(0)
		} else {
			fmt.Println("You do not have access to container:", containerName)
			showHelp(containerNames, customImageNames)
			os.Exit(1)
		}
	}

	if command == "stop" {
		if len(os.Args) < 3 {
			showHelp(containerNames, customImageNames)
			os.Exit(1)
		}

		// Check permitted name
		containerName := os.Args[2]
		if arrayContains(containerNames, containerName) {
			showStop(containerName)
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

		// Check custom image
		imageName := os.Args[2]
		if arrayContains(customImageNames, imageName) {
			showRunCustomImage(imageName, userid, userHome, network)
			os.Exit(0)
		} else {

			// Check valid image name
			if validateImageName(imageName) {
				showRunImage(imageName, userid, userHome, network)
				os.Exit(0)
			} else {
				fmt.Println("The image name is invalid")
				showHelp(containerNames, customImageNames)
				os.Exit(1)
			}

		}
	}

	if command == "exec" {
		if len(os.Args) < 4 {
			showHelp(containerNames, customImageNames)
			os.Exit(1)
		}

		// Check custom image
		imageName := os.Args[2]
		cmd := os.Args[3]
		arguments := os.Args[4:]
		if arrayContains(customImageNames, imageName) {
			// Get command and arguments
			showExecCustomImage(imageName, cmd, arguments, userid, userHome, network)
			os.Exit(0)
		} else {

			// Check valid image name
			if validateImageName(imageName) {
				showExecImage(imageName, cmd, arguments, userid, userHome, network)
				os.Exit(0)
			} else {
				fmt.Println("The image name is invalid")
				showHelp(containerNames, customImageNames)
				os.Exit(1)
			}

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
