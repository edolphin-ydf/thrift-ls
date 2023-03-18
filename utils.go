package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"go.lsp.dev/protocol"
)

func TokenToPosition(token antlr.Token) protocol.Position {
	return protocol.Position{
		Line:      uint32(token.GetLine() - 1),
		Character: uint32(token.GetColumn() - 1),
	}
}

func SymbolToRange(symbol antlr.Token) protocol.Range {
	return protocol.Range{
		Start: TokenToPosition(symbol),
		End:   protocol.Position{
			Line:      uint32(symbol.GetLine() - 1),
			Character: uint32(symbol.GetColumn() - 1 + len(symbol.GetText())),
		},
	}
}
