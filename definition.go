package main

import (
	"strings"

	parser "github.com/edolphin-ydf/thrift-ls/antlr/gen"
	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

func definition(currentFileURI protocol.DocumentURI, position protocol.Position) []protocol.Location {
	file, ok := WorkspaceInstance.Files[currentFileURI]
	if !ok {
		return nil
	}

	typeIdentifier := findTypeIdentifierAtPosition(file, position, false)

	if typeIdentifier == "" || typeIdentifier == "void" {
		return nil
	}

	return findDefinitionForType(file, typeIdentifier)
}

func findTypeIdentifierAtPosition(file *File, position protocol.Position, includeNextPos bool) string {
	visitor := &FieldTypeFindVisitor{
		BaseThriftVisitor: &parser.BaseThriftVisitor{},
		position:          position,
		includeNextPos:    includeNextPos,
	}
	visitor.Visit(file.Document)

	if visitor.FieldTypeCtx == nil {
		return ""
	}

	return visitor.FieldTypeCtx.GetText()
}

func findDefinitionForType(file *File, typeIdentifier string) []protocol.Location {
	var packageName string
	packageName, typeIdentifier = splitTypeIdentifier(typeIdentifier)

	if packageName == "" {
		// 找当前文件
		for _, s := range file.Structs {
			if s.Name.Name == typeIdentifier {
				symbol := s.Name.GetSymbol()
				return []protocol.Location{{
					URI:   file.URI,
					Range: SymbolToRange(symbol),
				}}
			}
		}

		for _, e := range file.Enums {
			if e.Name.Name == typeIdentifier {
				symbol := e.Name.GetSymbol()
				return []protocol.Location{{
					URI:   file.URI,
					Range: SymbolToRange(symbol),
				}}
			}
		}

		return nil
	}

	// 找include的文件
	for _, include := range file.Includes {
		// include 可能是  include "xxx/xxx.thrift"
		// 也可能是 include "../../xxx.thrift"
		// 而package name 只是 xxx
		if includeToPackageName(include) != packageName {
			continue
		}

		includeFileName := IncludeToFullPath(file.URI, include)
		includedFileURI := uri.File(includeFileName)
		includedFile, ok := WorkspaceInstance.Files[includedFileURI]
		if !ok {
			return nil
		}

		return findDefinitionForType(includedFile, typeIdentifier)
	}

	return nil
}

func splitTypeIdentifier(typeIdentifier string) (packageName string, typeName string) {
	lastIndexOfDot := strings.LastIndex(typeIdentifier, ".")

	if lastIndexOfDot == -1 {
		return "", typeIdentifier
	}

	return typeIdentifier[:lastIndexOfDot], typeIdentifier[lastIndexOfDot+1:]
}

func includeToPackageName(include string) string {
	include = strings.TrimSuffix(include, ".thrift")

	lastIndexOfSlash := strings.LastIndex(include, "/")
	if lastIndexOfSlash == -1 {
		return include
	}

	return include[lastIndexOfSlash+1:]
}
