//go:build android || freebsd || linux || netbsd || openbsd
// +build android freebsd linux netbsd openbsd

package webbrowser

// command to run when on android, freebsd, linux, netbsd, or openbsd
var openCommand *browserCommand = &browserCommand{"xdg-open", nil}
