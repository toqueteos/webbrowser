Webbrowser
==========

Webbrowser provides a simple API for opening web pages on your default browser. It's inspired on [Python's webbrowser](http://docs.python.org/3.3/library/webbrowser.html) package but lacks some of its features (open new window).

It just opens a webpage, most browsers will open it on a new tab.

Installation
============

Copy & Paste fans: `go get github.com/toqueteos/webbrowser`

Usage
=====

Just import the package (after you got it):

    `import "github.com/toqueteos/webbrowser"`

Then use the `Open` function.

    `Open("http://golang.org")`

Just in case, you have a very simple example on [`examples/simple.go`](https://github.com/toqueteos/webbrowser/blob/master/examples/simple.go).
