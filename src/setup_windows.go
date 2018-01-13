package client

import (
	"fmt"
	"os"
	"os/user"
)

func Setup() {
	// Get the current node serial number
	serial := GetSerialNumber()
	fmt.Printf("Current serial number is: %s\n", serial)

	user, _ := user.Current()
	fmt.Printf("User info: %s, %s\n", user.Username)
	fmt.Printf("User info: %s\n", user.Name)
	fmt.Printf("User info: %s\n", user.HomeDir)

	// Install chef if required
	_, err := InstallChef()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bootstrap failed: %s\n", err)
		os.Exit(1)
	}
}
