package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	parser "github.com/edolphin-ydf/thrift-lsp/antlr/gen"
	"go.lsp.dev/protocol"
)

const (
	FieldReqRequired = "required"
	FieldReqOptional = "optional"
)

type Name struct {
	antlr.TerminalNode
	Name string
}

type Field struct {
	*parser.FieldContext
	ID       int
	FieldReq string
	Name     Name
	Type     string
}

// define a thrift struct
type Struct struct {
	*parser.Struct_Context
	Name   Name
	Fields []*Field
}

type EnumField struct {
	*parser.Enum_fieldContext
	Name  Name
	Value int
}

type Enum struct {
	*parser.Enum_ruleContext
	Name   Name
	Fields []*EnumField
}

type Function struct {
	*parser.Function_Context
	FuncType string
	Name     Name
	Params   []*Field
}

type Service struct {
	*parser.ServiceContext
	Name  Name
	Funcs []*Function
}

type File struct {
	URI             protocol.DocumentURI
	Document        parser.IDocumentContext
	Includes        []string
	Structs         []*Struct
	Enums           []*Enum
	Services        []*Service
	DocumentVersion int32

	Text        string
	TextVersion int32
}
