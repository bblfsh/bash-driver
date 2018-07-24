package normalizer

import (
	. "gopkg.in/bblfsh/sdk.v2/uast/transformer"
	"gopkg.in/bblfsh/sdk.v2/uast"
)

var Preprocess = Transformers([][]Transformer{
	{
		// ResponseMetadata is a transform that trims response metadata from AST.
		//
		// https://godoc.org/gopkg.in/bblfsh/sdk.v2/uast#ResponseMetadata
		ResponseMetadata{
			TopLevelIsRootNode: false,
		},
	},
	{Mappings(Preprocessors...)},
}...)

var Normalize = Transformers([][]Transformer{

	{Mappings(Normalizers...)},
}...)

// Preprocessors is a block of AST preprocessing rules rules.
var Preprocessors = []Mapping{
	// ObjectToNode defines how to normalize common fields of native AST
	// (like node type, token, positional information).
	//
	// https://godoc.org/gopkg.in/bblfsh/sdk.v2/uast#ObjectToNode
	ObjectToNode{
		OffsetKey: "startOffset",
		EndOffsetKey: "endOffset",
	}.Mapping(),
}

func mapString(key string) Mapping {
	return MapSemantic(key, uast.String{}, MapObj(
		Obj{uast.KeyToken: Var("val")},
		Obj{
			"Value":  Var("val"),
			"Format": String(""),
		},
	))
}

func mapIdentifier(key string) Mapping {
	return MapSemantic(key, uast.Identifier{}, MapObj(
		Obj{uast.KeyToken: Var("val")},
		Obj{"Name": Var("val")},
	))
}
// Normalizers is the main block of normalization rules to convert native AST to semantic UAST.
var Normalizers = []Mapping{
	mapString("[Bash] unevaluated string (STRING2)"),
	mapString("[Bash] string"),
	mapString("[Bash] string content"),
	mapString("backquote shellcommand"),
	mapString("[Bash] File reference"),

	mapIdentifier("[Bash] word"),
	mapIdentifier("[Bash] variable"),
	mapIdentifier("[Bash] assignment_word"),

	MapSemantic("comment", uast.Comment{}, MapObj(
		Obj{
			uast.KeyToken: CommentText([2]string{}, "comm"),
		},
		CommentNode(false, "comm", nil),
	)),

	MapSemantic("[Bash] file reference", uast.RuntimeImport{}, MapObj(
		Obj{
			uast.KeyToken: Var("file"),
		},
		Obj{
			"Path": Var("file"),
		},
	)),

	MapSemantic("def", uast.FunctionGroup{}, MapObj(
		Obj{
		},
		Obj{
			"Nodes": Arr(
				UASTType(uast.Alias{}, Obj{
					"Name": UASTType(uast.Identifier{}, Obj{
						"Name": Var("name"),
					}),
					"Node": UASTType(uast.Function{}, Obj{
						"Type": UASTType(uast.FunctionType{},
							CasesObj("case_args",
								Obj{},
								Objs{
									{"Arguments": Var("args")},
									{"Arguments": Arr()},
								},
							)),
						"Body": UASTType(uast.Block{}, Obj{
							"Statements": Var("body"),
						}),
					}),
				}),
			),
		},
	)),

}
