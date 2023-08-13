// Package webbrowser provides a simple API for opening web pages on your
// default browser.
package webbrowser

import (
	"net/url"
	"os/exec"
)

func (b browserCommand) Command(s string) (*exec.Cmd, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	validUrl := ensureValidURL(u)

	b.args = append(b.args, validUrl)

	return exec.Command(b.cmd, b.args...), nil
}

func (b browserCommand) Open(s string) error {
	cmd, err := b.Command(s)
	if err != nil {
		return err
	}

	return cmd.Run()
}
