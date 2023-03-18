//
// visitor_test.go
// Copyright (C) 2023 bytedance <yangdongfeng.01@bytedance.com>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"encoding/json"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	parser "github.com/edolphin-ydf/thrift-lsp/antlr/gen"
)

func TestVisitor(t *testing.T) {
	input, _ := antlr.NewFileStream("./test/a.thrift")
	lexer := parser.NewThriftLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewThriftParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Document()

	visitor := &FileVisitor{}
	antlr.ParseTreeWalkerDefault.Walk(visitor, tree)

	d, _ := json.MarshalIndent(visitor.File, "", "\t")
	t.Log(string(d))
}
