package main

import (
	"fmt"

	"github.com/clburlison/bakeit/client"
)

func main() {
	config := client.Config()
	fmt.Printf("%s", config)

	client.Setup()
}
