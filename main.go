package main

import (
	"fmt"
	"os"

	"github.com/abs-lang/abs/repl"
)

var VERSION = "1.4.0"

// The ABS interpreter
func main() {
	args := os.Args
	if len(args) == 2 && args[1] == "--version" {
		fmt.Println(Version)
		return
	}
	// begin the REPL
	repl.BeginRepl(args, Version)
}
