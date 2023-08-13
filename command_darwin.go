//go:build darwin
// +build darwin

package webbrowser

// command to execute when running on a mac.
var openCommand *browserCommand = &browserCommand{"open", nil}
