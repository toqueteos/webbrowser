package webbrowser

import (
	"fmt"
	"net/url"
	"runtime"
	"strings"
)

// Open tries to open a URL in your default browser ensuring you have a display
// set up and not running this from SSH. NOTE: This may cause your program to
// hang until the browser process is closed in some OSes, see
// https://github.com/toqueteos/webbrowser/issues/4.
func Open(s string) (err error) {
	if openCommand == nil {
		return ErrNoCandidates
	}

	// Try to determine if there's a display available (only linux) and we
	// aren't on a terminal (all but windows).
	err = DisplayValid()
	if err != nil {
		return fmt.Errorf("webbrowser: tried to open %q, but %s", s, err.Error())
	}

	// Try all candidates
	for _, candidate := range Candidates {
		err := candidate.Open(s)
		if err == nil {
			return nil
		}
	}

	return ErrCantOpenBrowser
}

func ensureScheme(u *url.URL) {
	for _, s := range winSchemes {
		if u.Scheme == s {
			return
		}
	}
	u.Scheme = "http"
}

func ensureValidURL(u *url.URL) string {
	// Enforce a scheme (windows requires scheme to be set to work properly).
	ensureScheme(u)
	s := u.String()

	// Escape characters not allowed by cmd/bash
	switch runtime.GOOS {
	case "windows":
		s = strings.Replace(s, "&", `^&`, -1)
	}

	return s
}
