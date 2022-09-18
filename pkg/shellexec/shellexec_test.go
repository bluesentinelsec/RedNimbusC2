package shellexec

import "testing"

func TestShellExec(t *testing.T) {
	err := ExecShellCmd("hostname")
	if err != nil {
		t.Fatal(err)
	}
}
