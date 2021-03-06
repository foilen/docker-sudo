package main

import (
	"fmt"
)

func showHelp(containerNames []string, customImageNames []string) {
	fmt.Println("Commands:")
	fmt.Println("  help : Show this help")
	fmt.Println("  ps : Show the status of all the containers you can see")
	fmt.Println("  bash <containerName> : Go into a container you can see")
	fmt.Println("  logs <containerName> : View the logs of a container you can see")
	fmt.Println("  restart <containerName> : Restart a container you own")
	fmt.Println("  start <containerName> : Start a container you own")
	fmt.Println("  stop <containerName> : Stop a container you own")
	fmt.Println("  tails <containerName> : View the logs of a container you can see and tail it")
	fmt.Println("  run <customImageName or hubImageName> : Start an image with your home directory mounted and start bash")
	fmt.Println("  exec <customImageName or hubImageName> <command> [arg1] [argN]: Start an image with your home directory mounted and execute the command")
	fmt.Println("\nManageable containers by you:")
	for _, containerName := range containerNames {
		fmt.Println("  " + containerName)
	}
	fmt.Println("\nCustom images you can start:")
	for _, customImageName := range customImageNames {
		fmt.Println("  " + customImageName)
	}
}
