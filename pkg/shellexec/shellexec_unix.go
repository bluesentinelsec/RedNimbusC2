package shellexec

import (
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func ExecShellCmd(shellCmd string) error {
	log.Debug("executing command: \n", shellCmd)

	cmd := exec.Command("sh", "-c", shellCmd)

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", stdoutStderr)
	return err
}
