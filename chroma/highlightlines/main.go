package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func main() {
	lexer := lexers.Get("go")
	if lexer == nil {
		lexer = lexers.Fallback
	}
	style := styles.Get("monokai")
	if style == nil {
		style = styles.Fallback
	}
	highlightRanges := [][2]int{
		{4, 6},
	}
	formatter := html.New(
		html.TabWidth(3),
		html.WithLineNumbers(),
		// Note, it works well without the `WithClasses` option. But, with this
		// option the lines are not highlighted.
		html.WithClasses(),
		html.LineNumbersInTable(),
		html.HighlightLines(highlightRanges),
	)
	buf := &bytes.Buffer{}
	buf.WriteString("<!DOCTYPE html>\n<html>\n<head>\n<style>\n")
	if err := formatter.WriteCSS(buf, style); err != nil {
		log.Fatalf("%+v", err)
	}
	buf.WriteString("</style>\n</head>\n")
	iterator, err := lexer.Tokenise(nil, sourceCode)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	buf.WriteString("<body>\n")
	if err := formatter.Format(buf, style, iterator); err != nil {
		log.Fatalf("%+v", err)
	}
	buf.WriteString("</body>\n</html>")
	fmt.Println(buf.String())
}

const sourceCode = `package p

func f() {
	// this line is highlighted.
	// this one too.
	// as is this one.
	// But not this one.
}
`
