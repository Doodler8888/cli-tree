package zellij

import (
	// "os.Args"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)


var listSessions = &Z.Cmd{
	Name:    "sessions",
	Summary: "print all active sessions",
	Call: func(cmd *Z.Cmd, _ ...string) error {
		showSessions()
		return nil
	},
	Commands: []*Z.Cmd{
		help.Cmd,
		sessionKill,
		sessionsKill,
	},
}


var sessionAttach = &Z.Cmd{

	Name:    "attach",
	Summary: "attach to an active session",
	Call: func(cmd *Z.Cmd, args ...string) error {
		name := args[0]
	        attachSession(name)
		return nil
	},
	Commands: []*Z.Cmd{
		help.Cmd,
	},
}

var sessionKill = &Z.Cmd{

	Name:    "kill",
	Summary: "kill an active session",
	Call: func(cmd *Z.Cmd, args ...string) error {
		name := args[0]
	        killSession(name)
		return nil
	},
	Commands: []*Z.Cmd{
		help.Cmd,
	},
}

var sessionsKill = &Z.Cmd{

	Name:    "kill-all",
	Summary: "kill all active sessions",
	Call: func(cmd *Z.Cmd, _ ...string) error {
	        killSessions()
		return nil
	},
	Commands: []*Z.Cmd{
		help.Cmd,
	},
}




