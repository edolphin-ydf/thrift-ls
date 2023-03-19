package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"go.lsp.dev/protocol"
)

func TokenToPosition(token antlr.Token) protocol.Position {
	return protocol.Position{
		Line:      uint32(token.GetLine() - 1),
		Character: uint32(token.GetColumn()),
	}
}

func SymbolToRange(symbol antlr.Token) protocol.Range {
	return protocol.Range{
		Start: TokenToPosition(symbol),
		End: protocol.Position{
			Line:      uint32(symbol.GetLine() - 1),
			Character: uint32(symbol.GetColumn() + len(symbol.GetText())),
		},
	}
}

func PositionInRange(start, stop antlr.Token, position protocol.Position) bool {
	//logger.Sugar().Debug(start.GetLine()-1, start.GetColumn()-1, "|", stop.GetLine()-1, stop.GetColumn()-1, "||", position)

	return start.GetLine()-1 <= int(position.Line) && stop.GetLine()-1 >= int(position.Line) &&
		start.GetColumn()-1 <= int(position.Character) && stop.GetColumn()-1 >= int(position.Character)
}

func PositionInText(start antlr.Token, text string, position protocol.Position) bool {
	//logger.Sugar().Debug(start.GetLine()-1, start.GetColumn(), "|", start.GetColumn()-1 + len(text), "||", position)

	return start.GetLine()-1 == int(position.Line) &&
		start.GetColumn() <= int(position.Character) && start.GetColumn()-1+len(text) >= int(position.Character)
}

func PositionInOrAfterText(start antlr.Token, text string, position protocol.Position) bool {
	//logger.Sugar().Debug(start.GetLine()-1, start.GetColumn(), "|", start.GetColumn() + len(text), "||", position)

	return start.GetLine()-1 == int(position.Line) &&
		start.GetColumn() <= int(position.Character) && start.GetColumn()+len(text) >= int(position.Character)
}

func LogBaseParserRuleContext(ctx *antlr.BaseParserRuleContext) {
	start := ctx.GetStart()
	stop := ctx.GetStop()
	logger.Sugar().Debug(start.GetLine(), ",", start.GetColumn(), "|", stop.GetLine(), ",", stop.GetColumn())
}
