package docker

import (
	"fmt"
	"os/exec"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)


func NewCmd() *Z.Cmd {
	return &Z.Cmd{
		Name: "docker",
		Commands: []*Z.Cmd{
		 help.Cmd,
			// {
			// 	Name: "run",
			// },
			// {
			// 	Name: "pull",
			// },
			// {
			// 	Name: "push",
			// },
			{
				Name: "images",
				Summary: "List docker images",
				Call: func(cmd *Z.Cmd, _ ...string) error {
					out, err := exec.Command("docker", "images").Output()
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
			},
		},
	}
}
