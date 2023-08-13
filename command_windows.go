//go:build windows
// +build windows

package webbrowser

// command to execute when on a windows
var openCommand *browserCommand = &browserCommand{"cmd", []string{"/c", "start"}}
