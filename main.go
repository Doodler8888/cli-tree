package main

import (
	// "fmt"

	"github.com/Doodler8888/bonzai/pkg/docker"
	Z "github.com/rwxrob/bonzai/z"
	// "github.com/rwxrob/help"
)

func main() {
	rootCmd := &Z.Cmd{
		Name: "bonzai",
		Commands: []*Z.Cmd{  // So here "Commands: []*Z.Cmd{" i point to the &Z.Cmd, because if i don't do that, i will create a new object and it doesn't make  sense here from the efficiency point of view.
			docker.NewCmd(),
		},
	}

	rootCmd.Run()
}
