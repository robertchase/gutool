package cli

import (
    "flag"
    "fmt"
    "os"

    "github.com/robertchase/gutool/logic"
)

type appEnv struct {
    width int
    indent int
}


func CLI(args []string) int {
    var app appEnv
    if err := app.fromArgs(args); err != nil {
        return 2
    }
    if err := app.run(); err != nil {
        fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
        return 1
    }
    return 0
}

func (app *appEnv) fromArgs(args []string) error {
    fl := flag.NewFlagSet("ucol", flag.ContinueOnError)
    fl.IntVar(&app.width, "width", 80, "max width of line")
    fl.IntVar(&app.indent, "indent", -1, "left indent")
    if err := fl.Parse(args); err != nil {
        return err
    }
    return nil
}

func (app *appEnv) run() error {
    if err := logic.Upar(os.Stdin, os.Stdout, app.width, app.indent); err != nil {
        return err
    }
    return nil
}
