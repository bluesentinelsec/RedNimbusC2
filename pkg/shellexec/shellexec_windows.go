package shellexec

import (
	"fmt"
	"os/exec"
)

func ExecShellCmd(shellCmd string) error {
	cmd := exec.Command("cmd", "/c", shellCmd)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", stdoutStderr)
	return err
}
