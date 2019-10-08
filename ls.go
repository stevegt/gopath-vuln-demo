package main

import "fmt"

func main() {
	fmt.Println(`
	Move $GOPATH/bin to the end of $PATH to fix the security vulnerability described at
	https://github.com/stevegt/gopath-vuln-demo.  In the meantime, you can still get to 
	the system 'ls' command by running '/bin/ls' instead of 'ls'.
	`)
}
