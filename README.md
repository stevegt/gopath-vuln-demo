# $GOPATH/bin and $PATH vulnerability demo

**Disclaimer: DO NOT INSTALL this package.**

This is a short demo of one of the security vulnerabilities created when $GOPATH/bin is at the beginning of $PATH.  If you have $GOPATH/bin anywhere but at the end of your $PATH, then you probably want to fix that.  The decentralized nature of Go package distribution, coupled with the default behavior of `go get`, means that anyone with a github ID has the potential to silently supersede a standard system command with their own code.  A malicious actor might, for instance, replace `cat` with their own binary -- if it still acts like `cat`, it could also be doing other things every time you run it, and you may never notice.

## Demo

This package installs a Go binary named `$GOPATH/bin/ls`. If your $PATH includes $GOPATH/bin anywhere before /bin, then your /bin/ls command will be masked by `$GOPATH/bin/ls`, and instead of `ls` output you'll see this message:

```
Move $GOPATH/bin to the end of $PATH.
```
If you fix your $PATH as the message recommends, then `ls` will be back to normal.  You can optionally remove $GOPATH/bin/ls.

## Details

Take a look at `ls.go` in this package, and at the first line of `go.mod`.  All it takes for a malicious package to do bad things is **one** of these:

- have a go.mod `module` line ending in a Linux command -- that's all we're doing here
- have a package name ending in a standard Linux command, such as github.com/xxxx/cat or github.com/xxxx/ls
- have a package that is installed via ..., as in go get github.com/xxxx/foo/... If the package has a subdirectory named cat/cat.go, where cat.go contains a 'main' package and a 'main' function, then `go get` will drop a binary into $GOPATH/bin/cat.

