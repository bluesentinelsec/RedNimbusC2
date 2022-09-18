package main

import (
	"fmt"
	"os"

	"github.com/bluesentinelsec/rednimbusc2/pkg/cli"
)

func main() {
	fmt.Println("red operator main")
	cli.InvokeCLI(os.Args)
}
