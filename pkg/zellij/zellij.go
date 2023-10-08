package zellij


import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

func NewCmd() *Z.Cmd {
	return &Z.Cmd{
		Name: "zellij",
		Commands: []*Z.Cmd{
			help.Cmd,
			listSessions,
			sessionAttach,
		},
	}
}


