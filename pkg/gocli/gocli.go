package gocli


import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

func NewCmd() *Z.Cmd {
	return &Z.Cmd{
		Name: "go",
		Commands: []*Z.Cmd{
			help.Cmd,
			goInstall,
		},
	}
}


