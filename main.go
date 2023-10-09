package main

import (

	"github.com/Doodler8888/cli-tree/pkg/docker"
	"github.com/Doodler8888/cli-tree/pkg/tmux"
	"github.com/Doodler8888/cli-tree/pkg/gocli"
	"github.com/Doodler8888/cli-tree/pkg/zellij"
	"github.com/Doodler8888/cli-tree/pkg/tar"
	"github.com/Doodler8888/cli-tree/pkg/unzip"
	Z "github.com/rwxrob/bonzai/z"
	// "github.com/rwxrob/help"
)

func main() {
	rootCmd := &Z.Cmd{
		Name: "cli-tree",
		Commands: []*Z.Cmd{ 
			docker.NewCmd(),
			tmux.NewCmd(),
			gocli.NewCmd(),
			zellij.NewCmd(),
			tar.NewCmd(),
			unzip.NewCmd(),
		},
	}

	rootCmd.Run()
}


// Here "Commands: []*Z.Cmd{" i point to the &Z.Cmd, because if i don't do
// that, i will create a new object and it doesn't make  sense here from the
// efficiency point of view.
