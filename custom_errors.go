package webbrowser

import "errors"

var ErrCantOpenBrowser error = errors.New("webbrowser: can't open browser")
var ErrNoCandidates error = errors.New("webbrowser: no browser candidate found for your OS")
