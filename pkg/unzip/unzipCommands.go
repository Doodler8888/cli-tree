package unzip

import (
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)


var _extractToDirectory = &Z.Cmd{

	Name:    "directory",
	Summary: "extract a zip archive into a specific directory",
	Call: func(cmd *Z.Cmd, args ...string) error {
	  filename := args[0]
	  directoryname := args[1]
	  err := extractToDirectory(filename, directoryname)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	},
	Commands: []*Z.Cmd{
		help.Cmd,
	},
}

