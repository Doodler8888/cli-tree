package zellij

import (
	"fmt"
	// "os"
	"os/exec"
	// "path/filepath"
)

func attachSession(customName string) error {
	err := exec.Command("zellij", "a", customName).Run()
	if err != nil {
		print("Incorrect session name?")
	}

	return nil
}

func showSessions() {
	out, err := exec.Command("zellij", "list-sessions").Output()
	fmt.Print(string(out))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func killSessions() {
	err := exec.Command("zellij", "kill-all-sessions").Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func killSession(name string) {
	err := exec.Command("zellij", "kill-sessions", name).Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// The exec.Command function in Go creates a new Cmd object, which represents an
// external command that's ready to be run. However, creating a Cmd object
// doesn't actually run the command. To run the command, you need to call one of
// the methods on the Cmd object, such as Run or Output.
