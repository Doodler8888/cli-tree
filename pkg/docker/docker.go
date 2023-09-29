package docker

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

func NewCmd() *Z.Cmd {
	return &Z.Cmd{
		Name: "docker",
		Commands: []*Z.Cmd{
			help.Cmd,
			dockerImagesCmd,
			dockerImageDeleteCmd,
		},
	}
}
