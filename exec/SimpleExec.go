package exec

import (
	"errors"
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
		return "", errors.New(err.Error() + "\n" + content)
	}
	return content, nil
}
