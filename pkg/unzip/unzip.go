package unzip


import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"fmt"
)

func NewCmd() *Z.Cmd {
	return &Z.Cmd{
	Name:    "unzip",
	Summary: "extract a zip archive into current directory",
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
}


