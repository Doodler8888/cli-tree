package gocli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func installCustomBinary(customName string) error {
	// Build the binary with a custom name
	err := exec.Command("go", "build", "-o", customName).Run()
	if err != nil {
		return fmt.Errorf("failed to build: %v", err)
	}

	// Get GOPATH
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return fmt.Errorf("GOPATH is not set")
	}

	// Move the binary to $GOPATH/bin
	binPath := filepath.Join(gopath, "bin", customName)
	err = os.Rename(customName, binPath)
	if err != nil {
		return fmt.Errorf("failed to move binary: %v", err)
	}

	fmt.Printf("Successfully installed %s to %s\n", customName, binPath)
	return nil
}

func main() {
	customName := "your_custom_name_here" // Replace with the name you want
	err := installCustomBinary(customName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
