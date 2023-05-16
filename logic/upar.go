package logic

import (
	"bufio"
	"io"
	"strings"
)

type token struct {
	isNewParagraph bool
	isEof          bool
	value          string
}

func (tok *token) clear() {
	tok.isNewParagraph = false
	tok.isEof = false
	tok.value = ""
}

func newToken() *token {
	var t token
	t.clear()
	return &t
}

type tokenizer struct {
	input       *bufio.Scanner
	currentLine string
	line        []string
	isParagraph bool
	isEof       bool
}

func (t *tokenizer) nextLine() error {
	t.isEof = false
	t.isParagraph = false

	if t.input.Scan() {
		t.currentLine = t.input.Text()
        t.line = t.line[0:0]
		trimmed := strings.Trim(t.currentLine, " \t")
		if len(trimmed) == 0 {
			t.isParagraph = true
		} else {
            wordScanner := bufio.NewScanner(strings.NewReader(trimmed))
            wordScanner.Split(bufio.ScanWords)
            for wordScanner.Scan() {
                t.line = append(t.line, wordScanner.Text())
            }
		}
	} else {
		if err := t.input.Err(); err != nil {
			return err
		}
		t.isEof = true
	}
	return nil
}

func (t *tokenizer) nextNonBlankLine() error {
	for {
		if err := t.nextLine(); err != nil {
			return err
		}
		if t.isEof || !t.isParagraph {
			break
		}
	}
	return nil
}

func (t *tokenizer) nextToken(tok *token) error {
	tok.clear()
	if len(t.line) == 0 {
		if err := t.nextLine(); err != nil {
			return err
		}
		if t.isEof {
			tok.isEof = true
			return nil
		} else if t.isParagraph {
			tok.isNewParagraph = true
			return nil
		}
	}
	tok.value = t.line[0]
	t.line = t.line[1:]
	return nil
}

func newTokenizer(in io.Reader) *tokenizer {
	var t tokenizer
	t.input = bufio.NewScanner(in)
	t.line = []string{}
	return &t
}

func Upar(in io.Reader, out io.Writer, width int, indent int) error {
	source := newTokenizer(in)
	if err := source.nextNonBlankLine(); err != nil {
		return err
	}
	if source.isEof {
		return nil
	}

	var left string
	if indent < 0 {
		indent = len(source.currentLine) - len(strings.TrimLeft(source.currentLine, " \t"))
		left = source.currentLine[:indent]
	} else {
		left = strings.Repeat(" ", indent)
	}

	buffer := ""
	send := ""
	tok := newToken()
	writer := bufio.NewWriter(out)
	defer writer.Flush()
	newParagraph := false

	for tok.isEof == false {
		if err := source.nextToken(tok); err != nil {
			return err
		}

		if tok.isEof || tok.isNewParagraph {
			send = buffer
			buffer = ""
		} else if len(buffer) == 0 {
			if len(left+tok.value) > width {
				send = left + tok.value
			} else {
				buffer = left + tok.value
			}
		} else {
			if len(buffer+" "+tok.value) > width {
				send = buffer
				buffer = left + tok.value
			} else {
				buffer += " " + tok.value
			}
		}

		if len(send) > 0 {
			if newParagraph {
				send = "\n" + send
				newParagraph = false
			}
			if _, err := writer.WriteString(send + "\n"); err != nil {
				return err
			}
			send = ""
		}
		if tok.isNewParagraph {
			newParagraph = true
		}
	}
	return nil
}
