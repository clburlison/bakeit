package flight

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/clburlison/bakeit/src/config"
)

var (
	flightCmd  string
	flightArgs []string
	flightReq  bool
)

// Pre - Get command w/ args and whether it's required from config and store in variables
func Pre() {
	flightCmd = config.PreFlightCommand[runtime.GOOS]
	flightArgs = config.PreFlightArguments[runtime.GOOS]
	flightReq = config.PreFlightRequired[runtime.GOOS]

	err := runFlight("pre")
	if err != nil {
		os.Exit(1)
	}
}

// Post - Get command w/ args and whether it's required from config and store in variables
func Post() {
	flightCmd = config.PostFlightCommand[runtime.GOOS]
	flightArgs = config.PostFlightArguments[runtime.GOOS]
	flightReq = config.PostFlightRequired[runtime.GOOS]

	err := runFlight("post")
	if err != nil {
		os.Exit(1)
	}
}

func runFlight(timing string) error {
	if flightCmd == "" {
		fmt.Printf("%sflight command not configured\n", timing)
		return nil
	}

	// Run command with arguments
	fmt.Printf("Running %sflight command...\n", timing)
	cmd := exec.Command(flightCmd, flightArgs...)
	err := cmd.Run()

	if err != nil {
		fmt.Printf("%sflight command failed: %v\n", timing, err)
		if flightReq {
			// if error and required, log and exit
			fmt.Printf("Command required. Exiting...\n")
			return err
		}
		if !flightReq {
			// if error and not required, log and continue
			fmt.Printf("Command not required. Continuing...\n")
			return nil
		}
	}
	// if success. log and continue
	fmt.Printf("%sflight command completed successfully\n", timing)
	return nil
}
