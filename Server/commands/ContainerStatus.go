package commands

import (
	"os/exec"
)

func ContainerStatus() (string, error) {

	cmd := exec.Command("docker", "container", "ps")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
