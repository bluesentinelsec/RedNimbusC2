package shellexec

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

func ExecShellCmd(shellCmd string) error {
	log.Info("executing command: ", shellCmd)
	cmd := exec.Command("cmd", "/c", shellCmd)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", stdoutStderr)
	return err
}
