package zellij

import (
	"fmt"
	"os"
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
		fmt.Println("No active zellij sessions found.")
	}
}

func killSessions() {
    cmd := exec.Command("zellij", "kill-all-sessions")

    // Redirect the command's standard input, output and error to the shell
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Start()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    err = cmd.Wait()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}

func killSession(name string) {
	err := exec.Command("zellij", "kill-session", name).Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// The exec.Command function in Go creates a new Cmd object, which represents an
// external command that's ready to be run. However, creating a Cmd object
// doesn't actually run the command. To run the command, you need to call one of
// the methods on the Cmd object, such as Run or Output.

// The idea using 'cmd.Start()' is to probably initiate the command the way that
// the program doesn't care in that moment if the command completed or not. That
// way i avoid an error that will appear because of the incompetion.
// Then the prompt appears, and i targeting it by using 'cmd.Wait()'. Making the
// program possible to complete and close after the prompt. (it's just an
// assumption, i don't know really how right i'm here in realily)

