package tmux


import (
	// "fmt"
	// "os/exec"
	// "strings"
)

// func tmuxSessions() error {
// 	out, err := exec.Command("tmux", "list-sessions").Output()
// 	if err != nil {
// 		return nil
// 	}
// 	// Skip the header and split by newline
// 	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
// 	var groupedLines []string
// 	for i := 0; i < len(lines); i += 3 {
// 		groupedLines = append(groupedLines, lines[i]+"\n"+lines[i+1]+"\n"+lines[i+2])
// 	}
// 	return groupedLines, nil
// }


// The exec.Command function by itself only creates a new command object; it
// doesn't actually run the command. The Output() method is what runs the
// command, waits for it to complete, and collects its standard output
