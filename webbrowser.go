package webbrowser

import (
	"errors"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)

var (
	ErrCantOpen     = errors.New("webbrowser.Open: can't open webpage")
	ErrNoCandidates = errors.New("webbrowser.Open: no browser candidate found for your OS.")
)

// List of registered `Browser`s that will be tried with Open.
var candidates []Browser

type Browser interface {
	Open(string) error
}

// GenericBrowser just holds a command name and its arguments; the url will be
// appended as last arg. If you need to use string replacement for url define
// your own implementation.
type GenericBrowser struct {
	Cmd  string
	Args []string
}

func (gb GenericBrowser) Open(s string) error {
	u, err := url.Parse(s)
	if err != nil {
		return err
	}

	// Enforce a scheme
	if u.Scheme != "https" {
		u.Scheme = "http"
	}
	s = u.String()

	// Escape characters not allowed by cmd/bash
	switch runtime.GOOS {
	case "windows":
		s = strings.Replace(s, "&", `^&`, -1)
	default:
		s = strings.Replace(s, "&", `\&`, -1)
	}

	var cmd *exec.Cmd
	if gb.Args != nil {
		cmd = exec.Command(gb.Cmd, append(gb.Args, s)...)
	} else {
		cmd = exec.Command(gb.Cmd, s)
	}
	return cmd.Run()
}

// Open opens an URL on the first available candidate found.
func Open(s string) error {
	if len(candidates) == 0 {
		return ErrNoCandidates
	}

	for _, b := range candidates {
		err := b.Open(s)
		if err == nil {
			return nil
		}
	}

	return ErrCantOpen
}

// Register registers in the candidates list (append to end).
func Register(name Browser) {
	candidates = append(candidates, name)
}

// RegisterPrep registers in the candidates list (prepend to start).
func RegisterPrep(name Browser) {
	candidates = append([]Browser{name}, candidates...)
}

type args struct {
	cmd  string
	args []string
}

var osCommand = map[string]*args{
	"darwin":  &args{"open", nil},
	"freebsd": &args{"xdg-open", nil},
	"linux":   &args{"xdg-open", nil},
	"netbsd":  &args{"xdg-open", nil},
	"openbsd": &args{"xdg-open", nil}, // It may be open instead
	"windows": &args{"cmd", []string{"/c", "start"}},
}

func init() {
	// Register a generic browser, if any, for current OS.
	os, ok := osCommand[runtime.GOOS]
	if ok {
		Register(GenericBrowser{os.cmd, os.args})
	}
}
