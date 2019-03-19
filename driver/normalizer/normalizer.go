package normalizer

import (
	"gopkg.in/bblfsh/sdk.v2/uast"
	. "gopkg.in/bblfsh/sdk.v2/uast/transformer"
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
		OffsetKey:    "startOffset",
		EndOffsetKey: "endOffset",
	}.Mapping(),
}

func mapString(key string) Mapping {
	return MapSemantic(key, uast.String{}, MapObj(
		Obj{
			uast.KeyToken: Var("val"),
			"children":    Arr(),
		},
		Obj{
			"Value":  Var("val"),
			"Format": String(""),
		},
	))
}

func mapIdentifier(key string) Mapping {
	return MapSemantic(key, uast.Identifier{}, MapObj(
		Obj{
			uast.KeyToken: Var("val"),
			"children":    Arr(),
		},
		Obj{
			"Name": Var("val"),
		},
	))
}

var Normalizers = []Mapping{

	MapSemantic("function-def-element", uast.FunctionGroup{}, MapObj(
		Obj{
			"children": Arr(
				Obj{
					uast.KeyType: Var("_type_namedsymbol"),
					uast.KeyPos:  Var("_pos_namedsymbol"),
					"children": Arr(
						Obj{
							uast.KeyType: Var("_type_identifier"),
							uast.KeyPos:  Var("_pos_identifier"),
							"Name":       Var("name"),
						},
					),
				},
				Obj{
					uast.KeyType: Var("_type_groupelem"),
					uast.KeyPos:  Var("_pos_groupelem"),
					"children":   Var("body"),
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
						"Type": UASTType(uast.FunctionType{}, Obj{}),
						"Body": UASTType(uast.Block{}, Obj{
							"Statements": Var("body"),
						}),
					}),
				}),
			),
		},
	)),

	mapString("string_content"),
	mapString("string"),

	// replace "string" (aka string interpolation) with a single "string_content"
	// to a single uast:String node (already replace by previous transform)
	Map(
		Obj{
			uast.KeyType:  String("string"),
			uast.KeyToken: Any(),      // escaped string, don't need it in Semantic mode
			uast.KeyPos:   Var("pos"), // same as in the child node
			"children": One(
				Part("inner", Obj{
					uast.KeyType: String(uast.TypeOf(uast.String{})),
					uast.KeyPos:  Any(), // position without quotes; don't need it
				}),
			),
		},
		// TODO(dennwc): won't work for reversal
		Part("inner", Obj{
			uast.KeyType: String(uast.TypeOf(uast.String{})),
			uast.KeyPos:  Var("pos"), // position without quotes; don't need it
		}),
	),

	mapString("unevaluated_string2"),
	mapString("File_reference"),

	mapIdentifier("word"),
	mapIdentifier("variable"),
	mapIdentifier("assignment_word"),

	MapSemantic("Comment", uast.Comment{}, MapObj(
		Obj{
			uast.KeyToken: CommentText([2]string{"#", ""}, "comm"),
			"children":    Arr(),
		},
		CommentNode(false, "comm", nil),
	)),

	MapSemantic("file_reference", uast.RuntimeImport{}, MapObj(
		Obj{
			uast.KeyToken: Var("file"),
			"children":    Arr(),
		},
		Obj{
			"Path": Var("file"),
		},
	)),
}
