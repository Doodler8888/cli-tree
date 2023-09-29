package docker

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func getDockerImages() ([]string, error) {
	out, err := exec.Command("sh", "-c", `docker images | awk 'NR>1 {print $1 "\n       " $2 "\n       " $3 " " $4 " " $5 " " $6 " " $7}'`).Output()
	if err != nil {
		return nil, err
	}
	// Skip the header and split by newline
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	var groupedLines []string
	for i := 0; i < len(lines); i += 3 {
		groupedLines = append(groupedLines, lines[i]+"\n"+lines[i+1]+"\n"+lines[i+2])
	}
	return groupedLines, nil
}

func selectImages(images []string) ([]string, error) {
	var selectedImages []string
	prompt := &survey.MultiSelect{
		Message: "Select Docker images to delete:",
		Options: images,
	}
	err := survey.AskOne(prompt, &selectedImages)
	if err != nil { //
		fmt.Println(err)
		return nil, err //
	}

	return selectedImages, nil
}

func deleteDockerImages(images []string) error {
	for _, image := range images {
		err := exec.Command("docker", "rmi", image).Run()
		if err != nil { //
			fmt.Println(err)
			return err //
		}
	}

	return nil
}

// When you run docker images -q, the output will typically
// be a list of image IDs, one per line, like this:

// imageID1
// imageID2
// imageID3

// This entire output is captured in the out variable as a byte slice.
// After converting it to a string, it will look like
// "imageID1\nimageID2\nimageID3\n".

// The strings.Split(string(out), "\n") function will then split this string
// into a slice of strings, using the newline character \n as the delimiter.
// The resulting slice will look like this:

// ["imageID1", "imageID2", "imageID3", ""]

// Note that there's an empty string at the end because the original string had
// a trailing newline. Depending on your needs, you might want to remove this
// empty string from the slice before returning it.
