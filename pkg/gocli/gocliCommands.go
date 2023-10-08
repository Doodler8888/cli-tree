package gocli

import (
	"fmt"
	// "os.Args"
	"os/exec"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)


var installName = &Z.Cmd{
	Name:    "name",
	Summary: "go install command with an ability to choose a name for the binary",
	Call: func(cmd *Z.Cmd, args ...string) error {
		if len(args) < 1 {
			return fmt.Errorf("no binary name provided")
		}
		name := args[0]
		return installCustomBinary(name)
	},
	Commands: []*Z.Cmd{
		help.Cmd,
	},
}


var goInstall = &Z.Cmd{

	Name:    "install",
	Summary: "go insgall command with an ability to choose a name for the binary",
	Call: func(cmd *Z.Cmd, _ ...string) error {
		_, err := exec.Command("go", "install").Output()
		if err != nil {
			fmt.Println("Wrong command?", err)
			return err
		}
		// fmt.Println(string(out))
		return nil
	},
	Commands: []*Z.Cmd{
		help.Cmd,
		installName,
	},
}

