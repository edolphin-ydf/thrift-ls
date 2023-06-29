package main

import "go.lsp.dev/protocol"

func fileToSybolInformations(file *File) (res []interface{}) {
	for _, s := range file.Structs {
		res = append(res, &protocol.SymbolInformation{
			Name: s.Name.Name,
			Kind: protocol.SymbolKindStruct,
			Location: protocol.Location{
				URI:   file.URI,
				Range: SymbolToRange(s.Name.GetSymbol()),
			},
		})
	}

	for _, e := range file.Enums {
		res = append(res, &protocol.SymbolInformation{
			Name: e.Name.Name,
			Kind: protocol.SymbolKindStruct,
			Location: protocol.Location{
				URI:   file.URI,
				Range: SymbolToRange(e.Name.GetSymbol()),
			},
		})
	}

	for _, s := range file.Services {
		res = append(res, &protocol.SymbolInformation{
			Name: s.Name.Name,
			Kind: protocol.SymbolKindClass,
			Location: protocol.Location{
				URI:   file.URI,
				Range: SymbolToRange(s.Name.GetSymbol()),
			},
		})

		for _, f := range s.Funcs {
			res = append(res, &protocol.SymbolInformation{
				Name: f.Name.Name,
				Kind: protocol.SymbolKindFunction,
				Location: protocol.Location{
					URI:   file.URI,
					Range: SymbolToRange(f.Name.GetSymbol()),
				},
			})
		}
	}

	return
}
