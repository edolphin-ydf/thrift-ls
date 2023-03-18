package main

import (
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

	// service
	currentService *Service
	currentFunc    *Function
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

// EnterService is called when production service is entered.
func (s *FileVisitor) EnterService(ctx *parser.ServiceContext) {
	s.currentService = &Service{
		ServiceContext: ctx,
		Name: Name{
			TerminalNode: ctx.IDENTIFIER(0),
		},
	}

	identifier := ctx.IDENTIFIER(0)
	if identifier != nil {
		s.currentService.Name.Name = ctx.IDENTIFIER(0).GetText()
	}

	s.File.Services = append(s.File.Services, s.currentService)
}

// ExitService is called when production service is exited.
func (s *FileVisitor) ExitService(ctx *parser.ServiceContext) {
	s.currentService = nil
}

// EnterFunction_ is called when production function_ is entered.
func (s *FileVisitor) EnterFunction_(ctx *parser.Function_Context) {
	if s.currentService == nil {
		return
	}

	s.currentFunc = &Function{
		Function_Context: ctx,
		Name: Name{
			TerminalNode: ctx.IDENTIFIER(),
		},
	}

	identifier := ctx.IDENTIFIER()
	if identifier != nil {
		s.currentFunc.Name.Name = ctx.IDENTIFIER().GetText()
	}

	s.currentService.Funcs = append(s.currentService.Funcs, s.currentFunc)
}

// ExitFunction_ is called when production function_ is exited.
func (s *FileVisitor) ExitFunction_(ctx *parser.Function_Context) {
	s.currentFunc = nil
}

// EnterFunction_type is called when production function_type is entered.
func (s *FileVisitor) EnterFunction_type(ctx *parser.Function_typeContext) {
	if s.currentFunc == nil {
		return
	}

	if ctx.Field_type() != nil && ctx.Field_type().IDENTIFIER() != nil {
		s.currentFunc.FuncType = ctx.Field_type().IDENTIFIER().GetText()
	} else {
		s.currentFunc.FuncType = ctx.GetText()
	}
}

// ExitFunction_type is called when production function_type is exited.
func (s *FileVisitor) ExitFunction_type(ctx *parser.Function_typeContext) {
}

// EnterStruct_ is called when production struct_ is entered.
func (s *FileVisitor) EnterStruct_(ctx *parser.Struct_Context) {
	s.currentStruct = &Struct{
		Struct_Context: ctx,
		Name: Name{
			TerminalNode: ctx.IDENTIFIER(),
		},
	}

	identifier := ctx.IDENTIFIER()
	if identifier != nil {
		s.currentStruct.Name.Name = ctx.IDENTIFIER().GetText()
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
		},
	}

	identifier := ctx.IDENTIFIER()
	if identifier != nil {
		s.currentField.Name.Name = ctx.IDENTIFIER().GetText()
	}

	if s.currentStruct != nil {
		s.currentStruct.Fields = append(s.currentStruct.Fields, s.currentField)
	} else if s.currentFunc != nil {
		s.currentFunc.Params = append(s.currentFunc.Params, s.currentField)
	}
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

	identifier := ctx.IDENTIFIER()
	if identifier != nil {
		s.currentEnum.Name.Name = ctx.IDENTIFIER().GetText()
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

	var (
		name  string
		value int64
	)
	if ctx.Integer() != nil && ctx.Integer().INTEGER() != nil {
		value, _ = strconv.ParseInt(ctx.Integer().INTEGER().GetText(), 10, 64)
	}
	if ctx.IDENTIFIER() != nil {
		name = ctx.IDENTIFIER().GetText()
	}

	s.currentEnum.Fields = append(s.currentEnum.Fields, &EnumField{
		Enum_fieldContext: ctx,
		Name: Name{
			TerminalNode: ctx.IDENTIFIER(),
			Name:         name,
		},
		Value: int(value),
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
	visitor.File.URI = uri
	visitor.File.Document = tree

	antlr.ParseTreeWalkerDefault.Walk(visitor, tree)

	res = append(res, &visitor.File)
	res = append(res, ParseFileByIncludes(uri, visitor.File.Includes)...)

	return
}

func ParseFile(uri protocol.DocumentURI, text string) (res []*File) {
	input := antlr.NewInputStream(text)
	return ParseFileByStream(uri, input)
}

func ParseFileByFileName(fileName string) (res []*File) {
	uri := uri.File(fileName)
	file, ok := WorkspaceInstance.Files[uri]

	var input antlr.CharStream
	if ok {
		if file.DocumentVersion >= file.TextVersion {
			return nil
		}

		input = antlr.NewInputStream(file.Text)
	} else {
		var err error
		input, err = antlr.NewFileStream(fileName)
		if err != nil {
			logger.Sugar().Debug("ParseFileByPath open stream fail:", err)
			return
		}
	}

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

	return filepath.Clean(currentDir + "/" + include)
}

//region find node==============================================================
type NodeFindVisitor struct {
	parser.BaseThriftListener

	position protocol.Position

	FieldTypeCtx *parser.Field_typeContext
}

// EnterField_type is called when production field_type is entered.
func (s *NodeFindVisitor) EnterField_type(ctx *parser.Field_typeContext) {
	//logger.Sugar().Debugf("EnterField_type: %v", ctx.GetText())
	if PositionInOrAfterText(ctx.GetStart(), ctx.GetText(), s.position) {
		s.FieldTypeCtx = ctx
	}
}

func FindNodeByPosition(file *File, position protocol.Position) *parser.Field_typeContext {
	visitor := &NodeFindVisitor{
		position: position,
	}

	antlr.ParseTreeWalkerDefault.Walk(visitor, file.Document)

	return visitor.FieldTypeCtx
}

//endregion find node==============================================================
