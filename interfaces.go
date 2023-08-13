package webbrowser

import "os/exec"

type Browser interface {
	// Command returns a ready to be used Cmd that will open an URL.
	Command(string) (*exec.Cmd, error)
	// Open tries to open a URL in your default browser. NOTE: This may cause
	// your program to hang until the browser process is closed in some OSes,
	// see https://github.com/toqueteos/webbrowser/issues/4.
	Open(string) error
}
