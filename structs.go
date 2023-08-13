package webbrowser

// structure defining a command that will open up
// the webbrowser.
type browserCommand struct {
	// command to execute.
	cmd string
	// command arguments passed into exec.Command.
	args []string
}
