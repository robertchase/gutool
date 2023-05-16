package logic

import (
    "strings"
    "testing"
)

func TestNewToken(t *testing.T) {
    tok := newToken()
    if tok.isEof {
        t.Error("Excpecting isEof to be false")
    }
    if tok.isNewParagraph {
        t.Error("Excpecting isParagraph to be false")
    }
    if tok.value != "" {
        t.Error("Excepcting value to be an empty string")
    }
}

func TestNewTokenizer(t *testing.T) {

    tk := newTokenizer(strings.NewReader(""))
    if len(tk.line) != 0 {
        t.Error("Expecting zero length line")
    }
}

func TestNextNonBlankLine(t *testing.T) {
    tk := newTokenizer(strings.NewReader("\n\nHello   World\n\n"))
    tk.nextNonBlankLine()
    t.Log(tk.currentLine)

    if tk.currentLine != "Hello   World" {
        t.Error("Expecting 'Hello   World' to be the first line")
    }
    if len(tk.line) != 2 {
        t.Error("Expecting line length to be 2")
    }
}

func TestNextLine(t *testing.T) {
    tk := newTokenizer(strings.NewReader("Line1\n\nLine2\n"))

    t.Log(tk.nextLine())
    if tk.currentLine != "Line1" {
        t.Error("Expecting 'Line1' to be the first line")
    }
    t.Log(tk.nextLine())
    if !tk.isParagraph {
        t.Error("Expecting a blank line (new paragraph)")
    }
    t.Log(tk.nextLine())
    if tk.currentLine != "Line2" {
        t.Error("Expecting 'Line2' to be the second line")
    }
    t.Log(tk.nextLine())
    if !tk.isEof {
        t.Error("Expecting EOF")
    }
}

func TestNextToken(t *testing.T) {
    tk := newTokenizer(strings.NewReader("A B C\nD E"))
    tok := newToken()

    var tests = []string {"A", "B", "C", "D", "E"}

    for _, testTok := range tests {
        tk.nextToken(tok)
        if tok.value != testTok {
            t.Error("Expecting 'A'")
        }
    }
}
