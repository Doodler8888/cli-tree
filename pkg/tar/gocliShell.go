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


// In Go's fmt.Errorf function, %v is a format verb used to print the value of a
// variable. It's a placeholder that gets replaced with the actual value of the
// variable when the string is printed.
//
// In your example, fmt.Errorf("failed to build: %v", err), %v will be replaced
// with the value of err. This is useful for including dynamic content, such as
// error messages, in your strings.
