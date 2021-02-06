package osquery

import (
	"fmt"
	"os/exec"
)

func Print() {
	cmd := exec.Command("bash", "-c", "echo ~")
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Print(string(out))
}
