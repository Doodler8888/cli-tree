package tar

import (
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)


var _extract = &Z.Cmd{
	Name:    "extract",
	Summary: "extract a tar archive into current directory",
	Call: func(cmd *Z.Cmd, args ...string) error {
		if len(args) < 1 {
			return fmt.Errorf("no archive name provided")
		}
		name := args[0]
		return extract(name)
	},
	Commands: []*Z.Cmd{
		help.Cmd,
		_extractToDirectory,
	},
}


var _extractToDirectory = &Z.Cmd{

	Name:    "directory",
	Summary: "extract a tar archive into a specific directory",
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

