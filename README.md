Webbrowser
==========

Webbrowser provides a simple API for opening web pages on your default browser. It's inspired on [Python's webbrowser](http://docs.python.org/3.3/library/webbrowser.html) package but lacks some of its features (open new window).

It just opens a webpage, most browsers will open it on a new tab.

**Looking for alternatives/more features?** Try out https://github.com/pkg/browser it does what webbrowser does and more!

It is licensed under the MIT open source license, please see the [LICENSE.txt](https://github.com/toqueteos/webbrowser/blob/master/LICENSE.txt) file for more information.

Crossplatform support
=====================

The package is guaranteed to work on `windows`, `linux` and `darwin`. It also has default support for `freebsd`, `openbsd` and `netbsd` but these three have not been tested yet (that I'm aware of).

Installation
============

Copy & Paste fans: `go get github.com/toqueteos/webbrowser`

Usage
=====

Just import the package (after you got it):

    import "github.com/toqueteos/webbrowser"

Then use the `Open` function.

    webbrowser.Open("http://golang.org")

Just in case, you have a very simple example on [`examples/simple.go`](https://github.com/toqueteos/webbrowser/blob/master/examples/simple.go).

Extras
======

Miki Tebeka wrote a nicer version that wasn't on godoc.org when I did this. [Check it out!](https://bitbucket.org/tebeka/go-wise/src/d8db9bf5c4d1/desktop.go?at=default)
