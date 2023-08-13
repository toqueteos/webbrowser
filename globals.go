package webbrowser

// Candidates contains a list of registered `Browser`s that will be tried with Open.
var Candidates []Browser

// string array containing the three options for
// a link's scheme.
var winSchemes [3]string = [3]string{"https", "http", "file"}
