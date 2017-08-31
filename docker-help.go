package main

import (
	"fmt"
)

func showHelp(containerNames []string, imagesNames []string) {
	fmt.Println("Commands:")
	fmt.Println("  help : Show this help")
	fmt.Println("  ps : Show the status of all the containers you can see")
	fmt.Println("  bash <containerName> : Go into a container you can see")
	fmt.Println("  logs <containerName> : View the logs of a container you can see")
	fmt.Println("  tails <containerName> : View the logs of a container you can see and tail it")
	fmt.Println("  run <imageName> : Start an image with your home directory mounted")
	fmt.Println("\nManageable containers by you:")
	for _, containerName := range containerNames {
		fmt.Println("  " + containerName)
	}
	fmt.Println("\nImages you can start:")
	for _, imagesName := range imagesNames {
		fmt.Println("  " + imagesName)
	}
}
