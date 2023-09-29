package main

import (
	// "fmt"

	"github.com/Doodler8888/cli-tree/pkg/docker"
	"github.com/Doodler8888/cli-tree/pkg/tmux"
	"github.com/Doodler8888/cli-tree/pkg/gocli"
	Z "github.com/rwxrob/bonzai/z"
	// "github.com/rwxrob/help"
)

func main() {
	rootCmd := &Z.Cmd{
		Name: "cli-tree",
		Commands: []*Z.Cmd{ // So here "Commands: []*Z.Cmd{" i point to the &Z.Cmd, because if i don't do that, i will create a new object and it doesn't make  sense here from the efficiency point of view.
			docker.NewCmd(),
			tmux.NewCmd(),
			gocli.NewCmd(),
		},
	}

	rootCmd.Run()
}
