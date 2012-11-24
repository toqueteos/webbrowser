package webbrowser

import (
	"errors"
	"net/url"
	"os/exec"
)

var (
	ErrCantOpen     = errors.New("webbrowser.Open: can't open webpage")
	ErrNoCandidates = errors.New("webbrowser.Open: no browser candidate found for your OS.")
)

// List of registered `Browser`s that will be tried with Open.
var Candidates []Browser

type Browser interface {
	Open(string) error
}

// GenericBrowser just holds a command name (usually `/bin/bash` or `cmd`) and
// its arguments; the url will be appended as last arg. If you need to use
// string replacement for url define your own implementation.
type GenericBrowser struct {
	Cmd  string
	Args []string
}

func (gb GenericBrowser) Open(s string) error {
	u, err := url.Parse(s)
	if err != nil {
		return err
	}

	// Enforce a scheme so linux and darwin work properly
	if u.Scheme != "https" {
		u.Scheme = "http"
	}
	s = u.String()

	cmd := exec.Command(gb.Cmd, append(gb.Args, s)...)
	return cmd.Run()
}

// Open opens an URL on the first available candidate found.
func Open(s string) error {
	if len(Candidates) == 0 {
		return ErrNoCandidates
	}

	for _, b := range Candidates {
		err := b.Open(s)
		if err == nil {
			return nil
		}
	}

	return ErrCantOpen
}

// Register registers in the Candidates list (append to end).
func Register(name Browser) {
	Candidates = append(Candidates, name)
}

// RegisterPrep registers in the Candidates list (prepend to start).
func RegisterPrep(name Browser) {
	Candidates = append([]Browser{name}, Candidates...)
}
