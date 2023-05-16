package main

import (
    "os"
    "github.com/robertchase/gutool/cli/upar"
)

func main() {
    os.Exit(cli.CLI(os.Args[1:]))
}
