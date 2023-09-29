package tmux 


import (
	"fmt"
	"os/exec"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var tmuxSessions = &Z.Cmd{

	Name:    "sessions",
	Summary: "List tmux sessions",
	Call: func(cmd *Z.Cmd, _ ...string) error {
		out, err := exec.Command("tmux", "list-sessions").Output()
		if err != nil {
			fmt.Println("Wrong command?", err)
			return err
		}
		fmt.Println(string(out))
		return nil
	},
	Commands: []*Z.Cmd{
		help.Cmd,
	},
}
