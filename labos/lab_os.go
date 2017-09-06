package labos

import (
	"bytes"
	"os/exec"
)

func ExecuteCmd(command string, params []string) (string, error) {
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(command, params...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	output := out.String()

	return output, nil
}
