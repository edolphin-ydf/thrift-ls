package main

import (
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	parser "github.com/edolphin-ydf/thrift-lsp/antlr/gen"
	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

type FileVisitor struct {
	parser.BaseThriftListener

	File File

	// struct
	currentStruct *Struct
	currentField  *Field

	// enum
	currentEnum *Enum
}

func (s *FileVisitor) VisitTerminal(node antlr.TerminalNode) {
	//log.Println(node.GetText())
}

// EnterInclude_ is called when production include_ is entered.
func (s *FileVisitor) EnterInclude_(ctx *parser.Include_Context) {
	include := strings.Trim(ctx.LITERAL().GetText(), "\"'")
	s.File.Includes = append(s.File.Includes, include)
}

// ExitInclude_ is called when production include_ is exited.
func (s *FileVisitor) ExitInclude_(ctx *parser.Include_Context) {}

// EnterStruct_ is called when production struct_ is entered.
func (s *FileVisitor) EnterStruct_(ctx *parser.Struct_Context) {
	s.currentStruct = &Struct{
		Struct_Context: ctx,
		Name: Name{
			TerminalNode: ctx.IDENTIFIER(),
			Name:         ctx.IDENTIFIER().GetText(),
		},
	}
	s.File.Structs = append(s.File.Structs, s.currentStruct)
}

// ExitStruct_ is called when production struct_ is exited.
func (s *FileVisitor) ExitStruct_(ctx *parser.Struct_Context) {
	s.currentStruct = nil
}

// EnterField is called when production field is entered.
func (s *FileVisitor) EnterField(ctx *parser.FieldContext) {
	s.currentField = &Field{
		FieldContext: ctx,
		Name: Name{
			TerminalNode: ctx.IDENTIFIER(),
			Name:         ctx.IDENTIFIER().GetText(),
		},
	}
	s.currentStruct.Fields = append(s.currentStruct.Fields, s.currentField)
}

// ExitField is called when production field is exited.
func (s *FileVisitor) ExitField(ctx *parser.FieldContext) {
	s.currentField = nil
}

// EnterField_id is called when production field_id is entered.
func (s *FileVisitor) EnterField_id(ctx *parser.Field_idContext) {
	if s.currentField == nil {
		return
	}

	id, _ := strconv.ParseInt(ctx.Integer().INTEGER().GetText(), 10, 64)
	s.currentField.ID = int(id)
}

// ExitField_id is called when production field_id is exited.
func (s *FileVisitor) ExitField_id(ctx *parser.Field_idContext) {}

// EnterField_req is called when production field_req is entered.
func (s *FileVisitor) EnterField_req(ctx *parser.Field_reqContext) {
	if s.currentField == nil {
		return
	}

	s.currentField.FieldReq = ctx.GetText()
}

// ExitField_req is called when production field_req is exited.
func (s *FileVisitor) ExitField_req(ctx *parser.Field_reqContext) {}

// EnterField_type is called when production field_type is entered.
func (s *FileVisitor) EnterField_type(ctx *parser.Field_typeContext) {
	if s.currentField == nil {
		return
	}

	s.currentField.Type = ctx.GetText()
}

// ExitField_type is called when production field_type is exited.
func (s *FileVisitor) ExitField_type(ctx *parser.Field_typeContext) {}

// EnterEnum_rule is called when production enum_rule is entered.
func (s *FileVisitor) EnterEnum_rule(ctx *parser.Enum_ruleContext) {
	s.currentEnum = &Enum{
		Enum_ruleContext: ctx,
		Name: Name{
			TerminalNode: ctx.IDENTIFIER(),
			Name:         ctx.IDENTIFIER().GetText(),
		},
	}
	s.File.Enums = append(s.File.Enums, s.currentEnum)
}

// ExitEnum_rule is called when production enum_rule is exited.
func (s *FileVisitor) ExitEnum_rule(ctx *parser.Enum_ruleContext) {
	s.currentEnum = nil
}

// EnterEnum_field is called when production enum_field is entered.
func (s *FileVisitor) EnterEnum_field(ctx *parser.Enum_fieldContext) {
	if s.currentEnum == nil {
		return
	}

	value, _ := strconv.ParseInt(ctx.Integer().INTEGER().GetText(), 10, 64)
	s.currentEnum.Fields = append(s.currentEnum.Fields, &EnumField{
		Enum_fieldContext: ctx,
		Name:              Name{
			TerminalNode: ctx.IDENTIFIER(),
			Name:         ctx.IDENTIFIER().GetText(),
		},
		Value:             int(value),
	})
}

// ExitEnum_field is called when production enum_field is exited.
func (s *FileVisitor) ExitEnum_field(ctx *parser.Enum_fieldContext) {}


func ParseFileByStream(uri protocol.DocumentURI, input antlr.CharStream) (res []*File) {
	lexer := parser.NewThriftLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewThriftParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Document()

	visitor := &FileVisitor{}
	antlr.ParseTreeWalkerDefault.Walk(visitor, tree)
	visitor.File.URI = uri
	res = append(res, &visitor.File)

	res = append(res, ParseFileByIncludes(uri, visitor.File.Includes)...)

	return
}

func ParseFile(uri protocol.DocumentURI, text string) (res []*File) {
	input := antlr.NewInputStream(text)
	return ParseFileByStream(uri, input)
}

func ParseFileByFileName(fileName string) (res []*File) {
	input, err := antlr.NewFileStream(fileName)
	if err != nil {
		log.Println("ParseFileByPath open stream fail:", err)
		return
	}

	uri := uri.File(fileName)
	return ParseFileByStream(uri, input)
}

func ParseFileByIncludes(uri protocol.DocumentURI, includes []string) []*File {
	var res []*File
	for _, include := range includes {
		fileName := IncludeToFullPath(uri, include)
		res = append(res, ParseFileByFileName(fileName)...)
	}

	return res
}

func IncludeToFullPath(currentURI protocol.DocumentURI, include string) string {
	currentFile := currentURI.Filename()
	currentDir := filepath.Dir(currentFile)

	return filepath.Clean(currentDir+ "/" + include)
}
