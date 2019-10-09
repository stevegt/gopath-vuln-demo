package main

import (
	"fmt"
	"os"
)

func main() {
	txt := `

	This message is coming from the binary installed at '%s'.  

	You can still run the system 'ls' command by calling '/bin/ls' instead
	of 'ls'.  You can of course remove the '%s' binary as well.

	If you're seeing this message, it likely means that $GOPATH/bin is
	near the beginning of your $PATH.  This means you are exposed to the
	vulnerability described at https://github.com/stevegt/gopath-vuln-demo.

	You likely want to move $GOPATH/bin to the end of your $PATH, as shown
	at https://golang.org/doc/code.html#GOPATH.  

	Your $PATH is currently set to:
	
	%s

	`
	msg := fmt.Sprintf(txt, os.Args[0], os.Args[0], os.Getenv("PATH"))
	panic(msg)
}
