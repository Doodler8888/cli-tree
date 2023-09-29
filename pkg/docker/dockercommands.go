package docker

import (
	"fmt"
	"os/exec"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var dockerImageDeleteCmd = &Z.Cmd{
	Name:    "delete",
	Summary: "Delete docker images",
	Call: func(cmd *Z.Cmd, _ ...string) error {
		images, err := getDockerImages()
		if err != nil {     //
		 fmt.Println(err)
			return err  //
		}

		selectedImages, err := selectImages(images)
		if err != nil {     //
		 fmt.Println(err)
			return err  //
		}

		return deleteDockerImages(selectedImages)
	},
	Commands: []*Z.Cmd{
		help.Cmd,
	},
}

var dockerImagesCmd = &Z.Cmd{

	Name:    "images",
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
}
