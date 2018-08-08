package normalizer

import (
	. "gopkg.in/bblfsh/sdk.v2/uast/transformer"
	"gopkg.in/bblfsh/sdk.v2/uast"
)

var Preprocess = Transformers([][]Transformer{
	{
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

var Normalizers = []Mapping{

	MapSemantic("function-def-element", uast.FunctionGroup{}, MapObj(
		Obj{
			"children": Arr(
				Obj{
					uast.KeyType: Var("_type_namedsymbol"),
					uast.KeyPos: Var("_pos_namedsymbol"),
					"children": Arr(
						Obj{
							uast.KeyType: Var("_type_identifier"),
							uast.KeyPos: Var("_pos_identifier"),
							"Name": Var("name"),
						},
					),
				},
				Obj{
					uast.KeyType: Var("_type_groupelem"),
					uast.KeyPos: Var("_pos_groupelem"),
					"children": Var("body"),
				},
			),
		},
		Obj{
			"Nodes": Arr(
				UASTType(uast.Alias{}, Obj{
					"Name": UASTType(uast.Identifier{}, Obj{
						"Name": Var("name"),
					}),
					"Node": UASTType(uast.Function{}, Obj{
						"Type": UASTType(uast.FunctionType{}, Obj{
						}),
						"Body": UASTType(uast.Block{}, Obj{
							"Statements": Var("body"),
						}),
					}),
				}),
			),
		},
	)),

	mapString("unevaluated_string2"),
	mapString("string"),
	mapString("string_content"),
	mapString("backquote_shellcommand"),
	mapString("File_reference"),

	mapIdentifier("word"),
	mapIdentifier("variable"),
	mapIdentifier("assignment_word"),

	MapSemantic("comment", uast.Comment{}, MapObj(
		Obj{
			uast.KeyToken: CommentText([2]string{}, "comm"),
		},
		CommentNode(false, "comm", nil),
	)),

	MapSemantic("file_reference", uast.RuntimeImport{}, MapObj(
		Obj{
			uast.KeyToken: Var("file"),
		},
		Obj{
			"Path": Var("file"),
		},
	)),

}
