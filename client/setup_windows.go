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
	fmt.Printf("User info: %s, %s\n", user.Uid, user.Gid)
	fmt.Printf("User info: %s, %s\n", user.Username, user.Name)
	fmt.Printf("User info: %s\n", user.HomeDir)
	fmt.Printf("Windows is not supported yet!\n")
	os.Exit(1)
}
