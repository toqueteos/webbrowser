package webbrowser

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var (
	ErrCantOpenBrowser = errors.New("webbrowser: can't open browser")
	ErrNoCandidates    = errors.New("webbrowser: no browser candidate found for your OS")
)

// List of registered `Browser`s that will be tried with Open.
var Candidates []Browser

type Browser interface {
	Open(string) error
}

// Open opens an URL on the first available candidate found.
func Open(s string) (err error) {
	if len(Candidates) == 0 {
		return ErrNoCandidates
	}

	var lastError error
	for _, candidate := range Candidates {
		err := candidate.Open(s)
		if err == nil {
			return nil
		} else {
			lastError = err
		}
	}
	if lastError != nil {
		return lastError
	}

	// Try to determine if there's a display available (only linux) and we
	// aren't on a terminal (all but windows).
	switch runtime.GOOS {
	case "linux":
		// No display, no need to open a browser. Lynx users **MAY** have
		// something to say about this.
		if os.Getenv("DISPLAY") == "" {
			return fmt.Errorf("webbrowser: tried to open %q, no screen found", s)
		}
		fallthrough
	case "darwin":
		// Check SSH env vars.
		if os.Getenv("SSH_CLIENT") != "" || os.Getenv("SSH_TTY") != "" {
			return fmt.Errorf("webbrowser: tried to open %q, but you are running a shell session", s)
		}
	}

	return ErrCantOpenBrowser
}

func init() {
	// Register a generic browser, if any, for current OS.
	if os, ok := osCommand[runtime.GOOS]; ok {
		Candidates = append(Candidates, browserCommand{os.cmd, os.args})
	}
}

var (
	osCommand = map[string]*browserCommand{
		"darwin":  &browserCommand{"open", nil},
		"freebsd": &browserCommand{"xdg-open", nil},
		"linux":   &browserCommand{"xdg-open", nil},
		"netbsd":  &browserCommand{"xdg-open", nil},
		"openbsd": &browserCommand{"xdg-open", nil}, // It may be open instead
		"windows": &browserCommand{"cmd", []string{"/c", "start"}},
	}
)

type browserCommand struct {
	cmd  string
	args []string
}

func (b browserCommand) Open(s string) error {
	u, err := url.Parse(s)
	if err != nil {
		return err
	}

	// Enforce a scheme (windows requires scheme to be set to work properly).
	if u.Scheme != "https" {
		u.Scheme = "http"
	}
	s = u.String()

	// Escape characters not allowed by cmd/bash
	switch runtime.GOOS {
	case "windows":
		s = strings.Replace(s, "&", `^&`, -1)
	}

	var cmd *exec.Cmd
	if b.args != nil {
		cmd = exec.Command(b.cmd, append(b.args, s)...)
	} else {
		cmd = exec.Command(b.cmd, s)
	}

	return cmd.Run()
}
