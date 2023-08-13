//go:build !linux && !darwin
// +build !linux,!darwin

package webbrowser

func DisplayValid() (err error) {
	return nil
}
