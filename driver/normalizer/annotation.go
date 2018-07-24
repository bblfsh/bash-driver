package normalizer

import (
	"gopkg.in/bblfsh/sdk.v2/uast"
	"gopkg.in/bblfsh/sdk.v2/uast/role"
	. "gopkg.in/bblfsh/sdk.v2/uast/transformer"
	"gopkg.in/bblfsh/sdk.v2/uast/transformer/positioner"


)

var Native = Transformers([][]Transformer{
	{Mappings(Annotations...)},
	{RolesDedup()},
}...)

var Code = []CodeTransformer{
	positioner.NewFillLineColFromOffset(),
}

func annotateTypeToken(typ, token string, roles ...role.Role) Mapping {
	return AnnotateType(typ,
		FieldRoles{
			uast.KeyToken: {Op: String(token)},
		}, roles...)
}

// FIXME XXX remove comments prefix

var Annotations = []Mapping{
	AnnotateType("FILE", nil, role.File),
	AnnotateType("[Bash] Comment", nil, role.Comment, role.Noop),
	AnnotateType("[Bash] linefeed", nil, role.Whitespace, role.Noop),
	AnnotateType("WHITE_SPACE", nil, role.Whitespace, role.Noop),
	AnnotateType("[Bash] int literal", nil, role.Number, role.Literal, role.Primitive),
	AnnotateType("[Bash] unevaluated string (STRING2)", nil, role.Expression, role.String, role.Literal),
	AnnotateType("[Bash] string begin", nil, role.Expression, role.String, role.Block),
	AnnotateType("[Bash] string end", nil, role.Expression, role.String, role.Incomplete),
	AnnotateType("[Bash] string content", nil, role.Expression, role.String, role.Literal),
	AnnotateType("[Bash] string", nil, role.Expression, role.String, role.Literal, role.Primitive),
	AnnotateType("[Bash] let", nil, role.Statement, role.Incomplete),
	AnnotateType("[Bash] arithmetic command", nil, role.Expression, role.Arithmetic, role.Incomplete),
	AnnotateType("[Bash] arithmetic simple", nil, role.Expression, role.Arithmetic, role.Incomplete),
	AnnotateType("[Bash] arith ==", nil, role.Operator, role.Relational, role.Equal),
	AnnotateType("[Bash] =", nil, role.Operator, role.Assignment),
	AnnotateType("[Bash] arith !=", nil, role.Operator, role.Relational, role.Not, role.Equal),
	AnnotateType("[Bash] !=", nil, role.Operator, role.Relational, role.Not, role.Equal),
	AnnotateType("[Bash] <", nil, role.Operator, role.Relational, role.LessThan),
	AnnotateType("[Bash] <=", nil, role.Operator, role.Relational, role.LessThanOrEqual),
	AnnotateType("[Bash] >", nil, role.Operator, role.Relational, role.GreaterThan),
	AnnotateType("[Bash] >=", nil, role.Operator, role.Relational, role.GreaterThanOrEqual),
	AnnotateType("[Bash] ||", nil, role.Operator, role.Boolean, role.Or),
	AnnotateType("[Bash] |", nil, role.Operator, role.Boolean, role.Or),
	AnnotateType("[Bash] &&", nil, role.Operator, role.Boolean, role.And),
	AnnotateType("[Bash] &", nil, role.Operator, role.Boolean, role.And),
	AnnotateType("[Bash] cond_op !", nil, role.Operator, role.Boolean, role.Not),
	AnnotateType("[Bash] cond_op ==", nil, role.Expression, role.Relational, role.Operator, role.Equal),
	AnnotateType("[Bash] if", nil, role.Statement, role.If),
	AnnotateType("conditional shellcommand", nil, role.Expression, role.Condition),

	// These are more tokens that real semantic nodes, but they're in the AST tree
	// so we must tag them.
	AnnotateType("[Bash] [ for arithmetic", nil, role.Incomplete),
	AnnotateType("[Bash] ] for arithmetic", nil, role.Incomplete),
	AnnotateType("[Bash] [ (left conditional)", nil, role.Incomplete),
	AnnotateType("[Bash] ] (left conditional)", nil, role.Incomplete),
	AnnotateType("[Bash] ;", nil, role.Incomplete),
	AnnotateType("[Bash] {", nil, role.Incomplete),
	AnnotateType("[Bash] }", nil, role.Incomplete),
	AnnotateType("[Bash] :", nil, role.Incomplete),
	AnnotateType("[Bash] [ (left square)", nil, role.Incomplete),
	AnnotateType("[Bash] ] (right square)", nil, role.Incomplete),
	AnnotateType("[Bash] ;;", nil, role.Incomplete),
	AnnotateType("[Bash] [[ (left bracket)", nil, role.Incomplete),
	AnnotateType("[Bash] ]] (right bracket)", nil, role.Incomplete),
	AnnotateType("[Bash] (", nil, role.Incomplete),
	AnnotateType("[Bash] ((", nil, role.Incomplete),
	AnnotateType("[Bash] )", nil, role.Incomplete),
	AnnotateType("[Bash] ))", nil, role.Incomplete),
	AnnotateType("[Bash] backquote `", nil, role.Incomplete),
	AnnotateType("[Bash] $", nil, role.Incomplete),
	AnnotateType("ERROR_ELEMENT", nil, role.Incomplete),
	AnnotateType("[Bash] redirect list", nil, role.Incomplete),
	AnnotateType("[Bash] redirect element", nil, role.Incomplete),
	AnnotateType("[Bash] then", nil, role.Statement, role.If, role.Then),
	AnnotateType("logical block", nil, role.Block),

	// FIXME: no role in the uast for "end" or "end block/scope"
	AnnotateType("[Bash] fi", nil, role.Statement, role.Incomplete),
	AnnotateType("[Bash] do", nil, role.Statement, role.Block),
	AnnotateType("[Bash] in", nil, role.Expression, role.Binary, role.Operator, role.Relational, role.Contains),
	AnnotateType("var-use-element", nil, role.Expression, role.Variable, role.Identifier),
	AnnotateType("[Bash] variable", nil, role.Expression, role.Variable, role.Identifier),
	AnnotateType("[Bash] composed variable, like subshell", nil, role.Variable, role.Expression, role.Identifier, role.Incomplete),
	AnnotateType("[Bash] var substitution", nil, role.Variable, role.Expression, role.Identifier, role.Incomplete),
	AnnotateType("backquote shellcommand", nil, role.Expression, role.String, role.Literal, role.Call, role.Incomplete),
	AnnotateType("subshell shellcommand", nil, role.Expression, role.Call, role.Incomplete),
	AnnotateType("[Bash] pipeline command", nil, role.Expression, role.Call, role.Incomplete),
	AnnotateType("[Bash] generic bash command", nil, role.Expression, role.Incomplete),
	AnnotateType("[Bash] Shebang", nil, role.Comment, role.Pathname, role.Incomplete),
	AnnotateType("[Bash] shebang element", nil, role.Comment, role.Pathname, role.Incomplete),
	AnnotateType("[Bash] done", nil, role.Statement, role.Incomplete),
	AnnotateType("[Bash] &[0-9] filedescriptor", nil, role.Identifier, role.Receiver, role.Incomplete),
	AnnotateType("[Bash] word", nil, role.Expression, role.Identifier),
	AnnotateType("[Bash] combined word", nil, role.Expression, role.String, role.Identifier, role.Incomplete),
	AnnotateType("while loop", nil, role.Statement, role.While),
	AnnotateType("[Bash] while", nil, role.Statement, role.While),
	AnnotateType("until loop", nil, role.Statement, role.While, role.Incomplete),
	AnnotateType("[Bash] Parameter expansion operator '@@'", nil, role.Operator, role.Incomplete),
	AnnotateType("[Bash] Parameter expansion operator '##'", nil, role.Operator, role.Incomplete),
	AnnotateType("[Bash] Parameter expansion operator '%%'", nil, role.Operator, role.Incomplete),
	AnnotateType("[Bash] Parameter expansion operator '::'", nil, role.Operator, role.Incomplete),
	AnnotateType("[Bash] Parameter expansion operator '//'", nil, role.Operator, role.Incomplete),
	AnnotateType("[Bash] >>", nil, role.Operator, role.Incomplete),
	AnnotateType("[Bash] case", nil, role.Expression, role.Case),
	AnnotateType("[Bash] case pattern", nil, role.Expression, role.Case, role.Condition),
	AnnotateType("[Bash] case pattern list", nil, role.Case, role.Body, role.Block),
	AnnotateType("case pattern", nil, role.Statement, role.Switch),
	AnnotateType("[Bash] esac", nil, role.Statement, role.Incomplete),
	AnnotateType("let command", nil, role.Expression, role.Assignment, role.Incomplete),
	AnnotateType("[Bash] lazy LET expression", nil, role.Expression, role.Assignment, role.Incomplete),

	annotateTypeToken("[Bash] cond_op", "-eq", role.Operator, role.Relational, role.Equal),
	annotateTypeToken("[Bash] cond_op", "-ne", role.Operator, role.Relational, role.Not, role.Equal),
	annotateTypeToken("[Bash] cond_op", "-gt", role.Operator, role.Relational, role.GreaterThan),
	annotateTypeToken("[Bash] cond_op", "-ge", role.Operator, role.Relational, role.GreaterThanOrEqual),
	annotateTypeToken("[Bash] cond_op", "-lt", role.Operator, role.Relational, role.LessThan),
	annotateTypeToken("[Bash] cond_op", "-le", role.Operator, role.Relational, role.LessThanOrEqual),

	AnnotateType("simple-command", nil, role.Expression),
	AnnotateType("function-def-element", nil, role.Function, role.Declaration),
	AnnotateType("[Bash] function", nil, role.Function, role.Declaration, role.Block),
	AnnotateType("[Bash] named symbol", nil, role.Name, role.Identifier),
	AnnotateType("group element", nil, role.Body, role.Block),
	annotateTypeToken("[Bash] generic bash command", "break", role.Statement, role.Break),
	annotateTypeToken("[Bash] generic bash command", "continue", role.Statement, role.Continue),
	AnnotateType("include-command", nil, role.Statement, role.Import),
	AnnotateType("[Bash] File reference", nil, role.String, role.Import, role.Pathname, role.Identifier),
	AnnotateType("[Bash] =", nil, role.Operator, role.Assignment),

	AnnotateType("var-def-element", MapObj(Obj{
		"children": Arr(
			ObjectRoles("left"),
			ObjectRoles("operator"),
			ObjectRoles("right")),
	},
		Obj{
			"children": Arr(
				ObjectRoles("left", role.Assignment, role.Left),
				ObjectRoles("operator", role.Assignment, role.Operator),
				ObjectRoles("right", role.Right),
			),
		},
	), role.Expression, role.Assignment, role.Declaration),

	AnnotateType("if shellcommand", nil, role.Statement, role.If),

	AnnotateType("[Bash] if", MapObj(Obj{
		"children": Arr(ObjectRoles("condition")),
	}, Obj{
		"children": Arr(ObjectRoles("condition", role.If, role.Condition)),
	}), role.Statement, role.If),

	AnnotateType("[Bash] elif", MapObj(Obj{
		"children": Arr(ObjectRoles("condition")),
	}, Obj{
		"children": Arr(ObjectRoles("condition", role.Else, role.If, role.Condition)),
	}), role.Statement, role.Else, role.If),

	AnnotateType("[Bash] then", MapObj(Obj{
		"children": Arr(ObjectRoles("body")),
	}, Obj{
		"children": Arr(ObjectRoles("body", role.Then, role.Body, role.Block)),
	}), role.Then),

	AnnotateType("[Bash] else", MapObj(Obj{
		"children": Arr(ObjectRoles("body")),
	}, Obj{
		"children": Arr(ObjectRoles("body", role.Else, role.Body, role.Block)),
	}), role.Statement, role.Else),

	// for i; do something; done
	AnnotateType("for shellcommand", MapObj(Obj{
		"children": Arr(
			ObjectRoles("itervar"),
			ObjectRoles("body")),
	}, Obj{
		"children": Arr(
			ObjectRoles("itervar", role.For, role.Iterator, role.Expression),
			ObjectRoles("body", role.For, role.Body, role.Block)),
	}), role.For, role.Statement),

	// for i in a; do something; done
	AnnotateType("for shellcommand", MapObj(Obj{
		"children": Arr(
			ObjectRoles("itervar"),
			ObjectRoles("update"),
			ObjectRoles("body")),
	}, Obj{
		"children": Arr(
			ObjectRoles("itervar", role.For, role.Iterator, role.Expression),
			ObjectRoles("update", role.For, role.Update, role.Expression),
			ObjectRoles("body", role.For, role.Body, role.Block)),
	}), role.For, role.Statement),

	// for i; do something; done
	AnnotateType("while loop", MapObj(Obj{
		"children": Arr(
			ObjectRoles("condition"),
			ObjectRoles("body")),
	}, Obj{
		"children": Arr(
			ObjectRoles("condition", role.While, role.Expression, role.Condition),
			ObjectRoles("body", role.While, role.Body, role.Block)),
	}), role.While, role.Statement),

}
