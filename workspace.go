package main

import "go.lsp.dev/protocol"

type Workspace struct {
	RootURI protocol.DocumentURI

	Files map[protocol.DocumentURI]*File
}

var WorkspaceInstance = &Workspace{
	Files: map[protocol.URI]*File{},
}
