package tar


import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

func NewCmd() *Z.Cmd {
	return &Z.Cmd{
		Name: "tar",
		Commands: []*Z.Cmd{
			help.Cmd,
			_extract,
		},
	}
}


