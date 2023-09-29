package tmux


import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

func NewCmd() *Z.Cmd {
	return &Z.Cmd{
		Name: "tmux",
		Commands: []*Z.Cmd{
			help.Cmd,
			tmuxSessions,
		},
	}
}

