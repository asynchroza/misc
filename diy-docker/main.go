package main

import (
	"os/exec"
)

// unix doesn't support namespaces
func main() {

	// this command was not sandboxed
	cmd := exec.Command("scutil", "--set", "HostName", "container")
	output, err := cmd.Output()

	if err != nil {
		println(err.Error())
	} else {
		println(string(output))
	}
}
