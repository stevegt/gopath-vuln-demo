# $GOPATH/bin and $PATH vulnerability demo

**Disclaimer: DO NOT INSTALL this package.**

This is a short demo of one of the security vulnerabilities created
when $GOPATH/bin is at the beginning of $PATH.  If you have
$GOPATH/bin anywhere but at the end of your $PATH, then you probably
want to fix that.  The decentralized nature of Go package
distribution, coupled with the default behavior of `go get`, means
that anyone with e.g. a github ID has the potential to silently
supersede a standard system command with their own code.  A malicious
actor might, for instance, replace `cat` with their own binary.  If
the new binary still acts like `cat`, it could also be doing other
things every time you run it, and you may never notice.

## Details

All it takes for a malicious package to be able to do bad things is
**one** of these:

1. Have a go.mod `module` path ending in a string that matches the
   name of a common Linux command.  
    - That's all we're doing here. Take a look at the first line of `go.mod`.     
    - Note that the path you would pass to `go get` to install this
      package is `github.com/stevegt/gopath-vuln-demo`.  It looks
      innocent enough, but because of that `module` path, the package
      is able to install a binary named `ls`.
2. Have a package that is installed via ..., as in go get github.com/xxxx/foo/... 
    - If the package foo has a subdirectory named cat/cat.go, where
      cat.go contains a 'main' package and a 'main' function, then `go
      get github/xxxx/foo/...` will compile cat.go and install it as
      $GOPATH/bin/cat.  
    - This is so subtle that even the developer of the package may not
      realize they are masking a system command.
3. Have a package name ending in a standard Linux command, such as
   github.com/xxxx/cat or github.com/xxxx/ls
    - This one is at least more obvious.

## Demo

This package installs a Go binary named `$GOPATH/bin/ls`. If your
$PATH includes $GOPATH/bin anywhere before /bin, then your /bin/ls
command will be masked by `$GOPATH/bin/ls`, and instead of `ls` output
you'll see a message describing the vulnerability -- see main.go for
how the message is generated.

If you fix your $PATH as the message recommends, then `ls` will be
back to normal.  You can optionally remove $GOPATH/bin/ls.

