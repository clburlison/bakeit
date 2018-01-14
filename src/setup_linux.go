package setup

import (
	"fmt"
	"os"
	"os/user"
)

func Setup() {
	user, _ := user.Current()
	fmt.Printf("User info: %s, %s\n", user.Uid, user.Gid)
	fmt.Printf("User info: %s, %s\n", user.Username, user.Name)
	fmt.Printf("User info: %s\n", user.HomeDir)
	fmt.Printf("Linux is not supported yet!\n")
	os.Exit(1)
}
