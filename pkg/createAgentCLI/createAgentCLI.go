package createAgentCLI

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	log "github.com/sirupsen/logrus"
)

func InvokeCLI(args []string) {

	parser := argparse.NewParser(args[0], "TBD")

	// top level commands
	createAgentCmd := parser.NewCommand("create-agent", "TBD")

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}

	// setup console logging
	log.SetLevel(log.DebugLevel)

	if createAgentCmd.Happened() {
		log.Debug("invoke create-agent")
		return
	}
}
