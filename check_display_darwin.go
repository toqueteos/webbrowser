//go:build darwin
// +build darwin

package webbrowser

func DisplayValid() (err error) {
	// Check SSH env vars.
	if os.Getenv("SSH_CLIENT") != "" || os.Getenv("SSH_TTY") != "" {
		return errors.New("you are running a shell session")
	}
	return nil
}
