//go:build !windows && !darwin && !android && !freebsd && !linux && !netbsd && !openbsd
// +build !windows,!darwin,!android,!freebsd,!linux,!netbsd,!openbsd

package webbrowser

var openCommand *browserCommand = nil
