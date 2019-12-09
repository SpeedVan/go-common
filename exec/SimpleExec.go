package exec

import (
	"fmt"
	"os/exec"
)

// SimpleExec todo
func SimpleExec(command string, params ...string) (string, error) {
	cmd := exec.Command(command, params...)

	fmt.Println(cmd.Args)

	bs, err := cmd.CombinedOutput()
	content := string(bs)
	if err != nil {
		return content, err
	}
	return content, nil
}
