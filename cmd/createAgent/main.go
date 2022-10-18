package main

import (
	"os"

	"github.com/bluesentinelsec/rednimbusc2/pkg/createAgentCLI"
)

func main() {
	createAgentCLI.InvokeCLI(os.Args)
}
