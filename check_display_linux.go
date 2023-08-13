//go:build linux
// +build linux

package webbrowser

import (
	"errors"
	"os"
)

func DisplayValid() (err error) {
	// No display, no need to open a browser. Lynx users **MAY** have
	// something to say about this.
	if os.Getenv("DISPLAY") == "" {
		return errors.New("no screen found")
	}
	return nil
}
