package main

import (
	"context"
	"os"

	"go.lsp.dev/protocol"
)

type Server struct {
}

func (se *Server) Initialize(ctx context.Context, params *protocol.InitializeParams) (result *protocol.InitializeResult, err error) {
	WorkspaceInstance.RootURI = params.RootURI

	return &protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			DefinitionProvider: true,
			TextDocumentSync: &protocol.TextDocumentSyncOptions{
				OpenClose: true,
				Change:    protocol.TextDocumentSyncKindFull,
			},
			CompletionProvider: &protocol.CompletionOptions{
				ResolveProvider:   true,
				TriggerCharacters: []string{"."},
			},
		},
		ServerInfo: &protocol.ServerInfo{
			Name:    "thrift-lsp",
			Version: "0.0.1",
		},
	}, nil
}

func (se *Server) Initialized(ctx context.Context, params *protocol.InitializedParams) (err error) {
	logger.Sugar().Debug("Initialized")
	return nil
}

func (se *Server) Shutdown(ctx context.Context) (err error) {
	logger.Sugar().Debug("Shutdown")
	return nil
}

func (se *Server) Exit(ctx context.Context) (err error) {
	logger.Sugar().Debug("Exit")
	os.Exit(0)
	return nil
}

func (se *Server) WorkDoneProgressCancel(ctx context.Context, params *protocol.WorkDoneProgressCancelParams) (err error) {
	return nil
}

func (se *Server) LogTrace(ctx context.Context, params *protocol.LogTraceParams) (err error) {
	return nil
}

func (se *Server) SetTrace(ctx context.Context, params *protocol.SetTraceParams) (err error) {
	return nil
}

func (se *Server) CodeAction(ctx context.Context, params *protocol.CodeActionParams) (result []protocol.CodeAction, err error) {
	return nil, nil
}

func (se *Server) CodeLens(ctx context.Context, params *protocol.CodeLensParams) (result []protocol.CodeLens, err error) {
	return nil, nil
}

func (se *Server) CodeLensResolve(ctx context.Context, params *protocol.CodeLens) (result *protocol.CodeLens, err error) {
	return nil, nil
}

func (se *Server) ColorPresentation(ctx context.Context, params *protocol.ColorPresentationParams) (result []protocol.ColorPresentation, err error) {
	return nil, nil
}

func (se *Server) Completion(ctx context.Context, params *protocol.CompletionParams) (result *protocol.CompletionList, err error) {
	logger.Sugar().Debug("Completion")
	//logger.Sugar().Debug("CompletionContext", params)
	return nil, nil
}

func (se *Server) CompletionResolve(ctx context.Context, params *protocol.CompletionItem) (result *protocol.CompletionItem, err error) {
	return nil, nil
}

func (se *Server) Declaration(ctx context.Context, params *protocol.DeclarationParams) (result []protocol.Location, err error) {
	logger.Sugar().Debug("Declaration", params.TextDocument.URI, params.Position)
	return nil, nil
}

func (se *Server) Definition(ctx context.Context, params *protocol.DefinitionParams) (result []protocol.Location, err error) {
	logger.Sugar().Debug("Definition", params.TextDocument.URI, params.Position)
	return definition(params.TextDocument.URI, params.Position), nil
}

func (se *Server) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) (err error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Sugar().DPanic("DidChange", r)
		}
	}()

	logger.Sugar().Debug("DidChange", params.TextDocument.URI)
	if len(params.ContentChanges) == 0 {
		return nil
	}

	file, ok := WorkspaceInstance.Files[params.TextDocument.URI]
	if ok && file.Version >= params.TextDocument.Version {
		return nil
	}

	files := ParseFile(params.TextDocument.URI, params.ContentChanges[0].Text)
	if len(files) == 0 {
		return nil
	}
	files[0].Version = params.TextDocument.Version

	for _, f := range files {
		logger.Sugar().Debug("DidChange", f.URI, f.Version)
		WorkspaceInstance.Files[f.URI] = f
	}

	return nil
}

func (se *Server) DidChangeConfiguration(ctx context.Context, params *protocol.DidChangeConfigurationParams) (err error) {
	return nil
}

func (se *Server) DidChangeWatchedFiles(ctx context.Context, params *protocol.DidChangeWatchedFilesParams) (err error) {
	return nil
}

func (se *Server) DidChangeWorkspaceFolders(ctx context.Context, params *protocol.DidChangeWorkspaceFoldersParams) (err error) {
	return nil
}

func (se *Server) DidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) (err error) {
	logger.Sugar().Debug("DidClose", params.TextDocument.URI)
	return nil
}

func (se *Server) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) (err error) {
	logger.Sugar().Debug("DidOpen", params.TextDocument.URI)

	files := ParseFile(params.TextDocument.URI, params.TextDocument.Text)
	if len(files) == 0 {
		return nil
	}
	files[0].Version = params.TextDocument.Version

	for _, f := range files {
		logger.Sugar().Debug("DidOpen", f.URI, f.Version)
		WorkspaceInstance.Files[f.URI] = f
	}
	return nil
}

func (se *Server) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) (err error) {
	logger.Sugar().Debug("DidSave", params.TextDocument.URI)
	return nil
}

func (se *Server) DocumentColor(ctx context.Context, params *protocol.DocumentColorParams) (result []protocol.ColorInformation, err error) {
	return nil, nil
}

func (se *Server) DocumentHighlight(ctx context.Context, params *protocol.DocumentHighlightParams) (result []protocol.DocumentHighlight, err error) {
	return nil, nil
}

func (se *Server) DocumentLink(ctx context.Context, params *protocol.DocumentLinkParams) (result []protocol.DocumentLink, err error) {
	return nil, nil
}

func (se *Server) DocumentLinkResolve(ctx context.Context, params *protocol.DocumentLink) (result *protocol.DocumentLink, err error) {
	return nil, nil
}

func (se *Server) DocumentSymbol(ctx context.Context, params *protocol.DocumentSymbolParams) (result []interface{}, err error) {
	logger.Sugar().Debug("DocumentSymbol", params.TextDocument.URI)
	return nil, nil
}

func (se *Server) ExecuteCommand(ctx context.Context, params *protocol.ExecuteCommandParams) (result interface{}, err error) {
	return nil, nil
}

func (se *Server) FoldingRanges(ctx context.Context, params *protocol.FoldingRangeParams) (result []protocol.FoldingRange, err error) {
	return nil, nil
}

func (se *Server) Formatting(ctx context.Context, params *protocol.DocumentFormattingParams) (result []protocol.TextEdit, err error) {
	return nil, nil
}

func (se *Server) Hover(ctx context.Context, params *protocol.HoverParams) (result *protocol.Hover, err error) {
	return nil, nil
}

func (se *Server) Implementation(ctx context.Context, params *protocol.ImplementationParams) (result []protocol.Location, err error) {
	return nil, nil
}

func (se *Server) OnTypeFormatting(ctx context.Context, params *protocol.DocumentOnTypeFormattingParams) (result []protocol.TextEdit, err error) {
	return nil, nil
}

func (se *Server) PrepareRename(ctx context.Context, params *protocol.PrepareRenameParams) (result *protocol.Range, err error) {
	return nil, nil
}

func (se *Server) RangeFormatting(ctx context.Context, params *protocol.DocumentRangeFormattingParams) (result []protocol.TextEdit, err error) {
	return nil, nil
}

func (se *Server) References(ctx context.Context, params *protocol.ReferenceParams) (result []protocol.Location, err error) {
	return nil, nil
}

func (se *Server) Rename(ctx context.Context, params *protocol.RenameParams) (result *protocol.WorkspaceEdit, err error) {
	return nil, nil
}

func (se *Server) SignatureHelp(ctx context.Context, params *protocol.SignatureHelpParams) (result *protocol.SignatureHelp, err error) {
	return nil, nil
}

func (se *Server) Symbols(ctx context.Context, params *protocol.WorkspaceSymbolParams) (result []protocol.SymbolInformation, err error) {
	return nil, nil
}

func (se *Server) TypeDefinition(ctx context.Context, params *protocol.TypeDefinitionParams) (result []protocol.Location, err error) {
	logger.Sugar().Debug("TypeDefinition", params.TextDocument.URI, params.Position)
	return definition(params.TextDocument.URI, params.Position), nil
}

func (se *Server) WillSave(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (err error) {
	return nil
}

func (se *Server) WillSaveWaitUntil(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (result []protocol.TextEdit, err error) {
	return nil, nil
}

func (se *Server) ShowDocument(ctx context.Context, params *protocol.ShowDocumentParams) (result *protocol.ShowDocumentResult, err error) {
	return nil, nil
}

func (se *Server) WillCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (result *protocol.WorkspaceEdit, err error) {
	return nil, nil
}

func (se *Server) DidCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (err error) {
	return nil
}

func (se *Server) WillRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (result *protocol.WorkspaceEdit, err error) {
	return nil, nil
}

func (se *Server) DidRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (err error) {
	return nil
}

func (se *Server) WillDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (result *protocol.WorkspaceEdit, err error) {
	return nil, nil
}

func (se *Server) DidDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (err error) {
	return nil
}

func (se *Server) CodeLensRefresh(ctx context.Context) (err error) {
	return nil
}

func (se *Server) PrepareCallHierarchy(ctx context.Context, params *protocol.CallHierarchyPrepareParams) (result []protocol.CallHierarchyItem, err error) {
	return nil, nil
}

func (se *Server) IncomingCalls(ctx context.Context, params *protocol.CallHierarchyIncomingCallsParams) (result []protocol.CallHierarchyIncomingCall, err error) {
	return nil, nil
}

func (se *Server) OutgoingCalls(ctx context.Context, params *protocol.CallHierarchyOutgoingCallsParams) (result []protocol.CallHierarchyOutgoingCall, err error) {
	return nil, nil
}

func (se *Server) SemanticTokensFull(ctx context.Context, params *protocol.SemanticTokensParams) (result *protocol.SemanticTokens, err error) {
	return nil, nil
}

func (se *Server) SemanticTokensFullDelta(ctx context.Context, params *protocol.SemanticTokensDeltaParams) (result interface{}, err error) {
	return nil, nil
}

func (se *Server) SemanticTokensRange(ctx context.Context, params *protocol.SemanticTokensRangeParams) (result *protocol.SemanticTokens, err error) {
	return nil, nil
}

func (se *Server) SemanticTokensRefresh(ctx context.Context) (err error) {
	return nil
}

func (se *Server) LinkedEditingRange(ctx context.Context, params *protocol.LinkedEditingRangeParams) (result *protocol.LinkedEditingRanges, err error) {
	return nil, nil
}

func (se *Server) Moniker(ctx context.Context, params *protocol.MonikerParams) (result []protocol.Moniker, err error) {
	return nil, nil
}

func (se *Server) Request(ctx context.Context, method string, params interface{}) (result interface{}, err error) {
	return nil, nil
}
