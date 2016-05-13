package main

import (
	"flag"
	"log"

	"github.com/kr/pretty"
	"github.com/mewspring/poc/empty/lexer"
	"github.com/mewspring/poc/empty/parser"
)

func main() {
	flag.Parse()
	for _, path := range flag.Args() {
		s, err := lexer.NewLexerFile(path)
		if err != nil {
			log.Fatal(err)
		}
		p := parser.NewParser()
		file, err := p.Parse(s)
		if err != nil {
			log.Fatal(err)
		}
		pretty.Println(file)
	}
}
