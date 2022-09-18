package cli

import "testing"

func TestInvokeCLI(t *testing.T) {

	// just confirm this package compiles
	args := []string{"prog", "set-task", "--cmd", "test"}
	InvokeCLI(args)
}
