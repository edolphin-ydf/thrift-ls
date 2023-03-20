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

func GetNodeRange(ctx antlr.ParserRuleContext) protocol.Range {
	startToken := ctx.GetStart()
	stopToken := ctx.GetStop()

	stopLine := stopToken.GetLine()
	stopCol := stopToken.GetColumn()
	if stopLine == startToken.GetLine() && stopCol == startToken.GetColumn() {
		stopCol += len(ctx.GetText())
	}

	return protocol.Range{
		Start: TokenToPosition(startToken), // 开始不-1
		End: protocol.Position{
			Line:      uint32(stopLine) - 1,
			Character: uint32(stopCol) - 1, // 结尾-1
		},
	}
}

func PositionInRange(r protocol.Range, position protocol.Position) bool {
	logger.Sugar().Debug(r, "|", position)

	start := r.Start
	stop := r.End

	// 行在范围外
	if start.Line > position.Line || stop.Line < position.Line {
		return false
	}

	if start.Line != stop.Line {
		// 开始和结束不在同一行
		if start.Line == position.Line {
			// 在开始行
			return start.Character <= position.Character
		} else if stop.Line == position.Line {
			// 在结束行
			return stop.Line >= position.Character
		} else {
			// 在中间行
			return true
		}
	} else {
		// 开始和结束在同一行, 只需要判断列是否在范围内
		return start.Character <= position.Character && stop.Character >= position.Character
	}
}

func PositionInOrAfterText(start antlr.Token, text string, position protocol.Position) bool {
	//logger.Sugar().Debug(start.GetLine()-1, start.GetColumn(), "|", start.GetColumn() + len(text), "||", position)

	return start.GetLine()-1 == int(position.Line) &&
		start.GetColumn() <= int(position.Character) && start.GetColumn()+len(text) >= int(position.Character)
}
