// Package main
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/pravj/rugo/parser"
	"github.com/pravj/rugo/scanner"
	"github.com/pravj/rugo/utils"
	"github.com/pravj/rugo/transpiler"
	"github.com/pravj/rugo/program"
)

func main() {
	// read the S-expression file
	content, err := ioutil.ReadFile("expr.s")
	utils.CheckError(err)

	// scan the input string
	s := scanner.NewScanner(string(content))
	s.ScanTokens()
	fmt.Println(s.Tokens)

	// token output
	//fmt.Println(string(content))
	//fmt.Println(s.Tokens)

	// parse the input
	root := parser.NewParser(s.Tokens).Parse()
	fmt.Println(root)

	// generate a Go program AST from the tokens
	p := program.NewProgram("main", root)
	transpiler.Transpile(root, p)

	fmt.Println(p.String())
}
