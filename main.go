// Package main
package main

import (
	"fmt"
	//"reflect"
	"io/ioutil"

	"github.com/pravj/rugo/parser"
	"github.com/pravj/rugo/scanner"
	"github.com/pravj/rugo/utils"
)

func main() {
	// read the S-expression file
	content, err := ioutil.ReadFile("expr.s")
	utils.CheckError(err)

	// scan the input string
	s := scanner.NewScanner(string(content))
	s.ScanTokens()

	// token output
	//fmt.Println(string(content))
	//fmt.Println(s.Tokens)

	p := parser.NewParser(s.Tokens)
	root := p.Parse()

	fmt.Println(root)
}
