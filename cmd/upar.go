package main

import (
	"github.com/robertchase/gutool/cli/upar"
	"os"
)

func main() {
	os.Exit(cli.CLI(os.Args[1:]))
}
