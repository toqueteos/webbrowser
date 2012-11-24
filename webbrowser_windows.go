package webbrowser

func init() {
	Register(GenericBrowser{"cmd", []string{"/c", "start"}})
}
