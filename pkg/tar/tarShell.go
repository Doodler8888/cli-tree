package tar

import (
	"fmt"
	"os/exec"
)

func extract(name string) error {
	err := exec.Command("tar", "-xvf", name).Run()
	if err != nil {
	  return fmt.Errorf("failed to extract %s: %v", name, err)
	}
	return nil
}

func extractToDirectory(filename string, directoryname string) error {
	err := exec.Command("tar", "-xvf", filename, "-C", directoryname).Run()
	if err != nil {
	  return fmt.Errorf("failed to extract %s: %v", filename, err)
	}
	return nil
}


// In Go's fmt.Errorf function, %v is a format verb used to print the value of a
// variable. It's a placeholder that gets replaced with the actual value of the
// variable when the string is printed.
//
// In your example, fmt.Errorf("failed to build: %v", err), %v will be replaced
// with the value of err. This is useful for including dynamic content, such as
// error messages, in your strings.
