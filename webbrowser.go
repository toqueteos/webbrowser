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

//
var Candidates []Browser

// Browser
type Browser interface {
	Open(string) error
}

// GenericBrowser
type GenericBrowser struct {
	cmd  string
	args []string
}

func (gb GenericBrowser) Open(s string) error {
	u, err := url.Parse(s)
	if err != nil {
		return err
	}

	u.Scheme = "http"
	s = u.String()

	cmd := exec.Command(gb.cmd, append(gb.args, s)...)
	return cmd.Run()
}

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

// Register a browser connector and, optionally, connection.
func Register(name Browser) {
	// Append
	Candidates = append(Candidates, name)
	// Prepend
	// Candidates = append([]Browser{name}, Candidates...)
}
