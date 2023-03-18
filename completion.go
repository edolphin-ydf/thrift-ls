package main

import (
	"strings"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

func completion(file *File, textBeforeCursor string) (res []protocol.CompletionItem) {
	packageName, partFieldName := splitTypeIdentifier(textBeforeCursor)
	logger.Sugar().Debug("packageName: ", packageName, " partFieldName: ", partFieldName)

	if packageName == "" {
		// include
		for _, include := range file.Includes {
			includePackage := includeToPackageName(include)
			if strings.HasPrefix(includePackage, partFieldName) {
				res = append(res, includeToCompletionItem(includePackage, partFieldName))
			}
		}

		// struct
		for _, s := range file.Structs {
			if strings.HasPrefix(s.Name.Name, partFieldName) {
				res = append(res, structToCompletionItem(s, partFieldName))
			}
		}

		// enum
		for _, e := range file.Enums {
			if strings.HasPrefix(e.Name.Name, partFieldName) {
				res = append(res, enumToCompletionItem(e, partFieldName))
			}
		}

		return
	}

	for _, include := range file.Includes {
		includePackage := includeToPackageName(include)
		if includePackage != packageName {
			if strings.HasPrefix(includePackage, packageName) {
				res = append(res, includeToCompletionItem(includePackage, partFieldName))
			}
			continue
		}

		logger.Sugar().Debug("includePackage: ", includePackage)
		includedFileName := IncludeToFullPath(file.URI, include)
		includedFileURI := uri.File(includedFileName)
		includedFile, ok := WorkspaceInstance.Files[includedFileURI]
		if !ok {
			return
		}

		return completion(includedFile, partFieldName)
	}

	return
}

func includeToCompletionItem(include string, prefix string) protocol.CompletionItem {
	return protocol.CompletionItem{
		Label:  include,
		Detail: "include",
		Kind:   protocol.CompletionItemKindModule,
	}
}

func structToCompletionItem(s *Struct, prefix string) protocol.CompletionItem {
	return protocol.CompletionItem{
		Label:  s.Name.Name,
		Detail: "struct",
		Kind:   protocol.CompletionItemKindClass,
	}
}

func enumToCompletionItem(e *Enum, prefix string) protocol.CompletionItem {
	return protocol.CompletionItem{
		Label:  e.Name.Name,
		Detail: "enum",
		Kind:   protocol.CompletionItemKindEnum,
	}
}
