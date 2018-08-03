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
	AnnotateType("Comment", nil, role.Comment, role.Noop),
	AnnotateType("int literal", nil, role.Number, role.Literal, role.Primitive),
	AnnotateType("unevaluated string (STRING2)", nil, role.Expression, role.String, role.Literal),
	AnnotateType("string content", nil, role.Expression, role.String, role.Literal),
	AnnotateType("string", nil, role.Expression, role.String, role.Literal, role.Primitive),
	AnnotateType("let", nil, role.Statement, role.Incomplete),
	AnnotateType("arithmetic command", nil, role.Expression, role.Arithmetic, role.Incomplete),
	AnnotateType("arithmetic simple", nil, role.Expression, role.Arithmetic, role.Incomplete),
	AnnotateType("arith ==", nil, role.Operator, role.Relational, role.Equal),
	AnnotateType("arith <", nil, role.Operator, role.Arithmetic, role.LeftShift),
	AnnotateType("arith >", nil, role.Operator, role.Arithmetic, role.RightShift),
	AnnotateType("=", nil, role.Operator, role.Assignment),
	AnnotateType("arith !=", nil, role.Operator, role.Relational, role.Not, role.Equal),
	AnnotateType("!=", nil, role.Operator, role.Relational, role.Not, role.Equal),
	AnnotateType("<", nil, role.Operator, role.Relational, role.LessThan),
	AnnotateType("<=", nil, role.Operator, role.Relational, role.LessThanOrEqual),
	AnnotateType(">", nil, role.Operator, role.Relational, role.GreaterThan),
	AnnotateType(">=", nil, role.Operator, role.Relational, role.GreaterThanOrEqual),
	AnnotateType("||", nil, role.Operator, role.Boolean, role.Or),
	AnnotateType("|", nil, role.Operator, role.Boolean, role.Or),
	AnnotateType("&&", nil, role.Operator, role.Boolean, role.And),
	AnnotateType("&", nil, role.Operator, role.Boolean, role.And),
	AnnotateType("cond_op !", nil, role.Operator, role.Boolean, role.Not),
	AnnotateType("cond_op ==", nil, role.Expression, role.Relational, role.Operator, role.Equal),
	AnnotateType("conditional shellcommand", nil, role.Expression, role.Condition),

	// These are more tokens that real semantic nodes, but they're in the AST tree
	// so we must tag them.
	AnnotateType("[ for arithmetic", nil, role.Incomplete),
	AnnotateType("] for arithmetic", nil, role.Incomplete),
	AnnotateType(":", nil, role.Incomplete),
	AnnotateType("[ (left square)", nil, role.Incomplete),
	AnnotateType("] (right square)", nil, role.Incomplete),
	AnnotateType(";;", nil, role.Incomplete),
	AnnotateType("[[ (left bracket)", nil, role.Incomplete),
	AnnotateType("]] (right bracket)", nil, role.Incomplete),
	AnnotateType("((", nil, role.Incomplete),
	AnnotateType("))", nil, role.Incomplete),
	AnnotateType("backquote `", nil, role.Incomplete),
	AnnotateType("$", nil, role.Incomplete),
	AnnotateType("ERROR_ELEMENT", nil, role.Incomplete),
	AnnotateType("redirect list", nil, role.Incomplete),
	AnnotateType("redirect element", nil, role.Incomplete),
	AnnotateType("then", nil, role.Statement, role.If, role.Then),
	AnnotateType("logical block", nil, role.Block),

	// FIXME: no role in the uast for "end" or "end block/scope"
	AnnotateType("in", nil, role.Expression, role.Binary, role.Operator, role.Relational, role.Contains),
	AnnotateType("var-use-element", nil, role.Expression, role.Variable, role.Identifier),
	AnnotateType("variable", nil, role.Expression, role.Variable, role.Identifier),
	AnnotateType("composed variable, like subshell", nil, role.Variable, role.Expression, role.Identifier, role.Incomplete),
	AnnotateType("var substitution", nil, role.Variable, role.Expression, role.Identifier, role.Incomplete),
	AnnotateType("backquote shellcommand", nil, role.Expression, role.String, role.Literal, role.Call, role.Incomplete),
	AnnotateType("subshell shellcommand", nil, role.Expression, role.Call, role.Incomplete),
	AnnotateType("pipeline command", nil, role.Expression, role.Call, role.Incomplete),
	AnnotateType("generic bash command", nil, role.Expression, role.Incomplete),
	AnnotateType("Shebang", nil, role.Comment, role.Pathname, role.Incomplete),
	AnnotateType("shebang element", nil, role.Comment, role.Pathname, role.Incomplete),
	AnnotateType("&[0-9] filedescriptor", nil, role.Identifier, role.Receiver, role.Incomplete),
	AnnotateType("word", nil, role.Expression, role.Identifier),
	AnnotateType("combined word", nil, role.Expression, role.Identifier, role.Incomplete),
	AnnotateType("while loop", nil, role.Statement, role.While),
	AnnotateType("until loop", nil, role.Statement, role.While, role.Incomplete),
	AnnotateType("Parameter expansion operator '@@'", nil, role.Operator, role.Incomplete),
	AnnotateType("Parameter expansion operator '@'", nil, role.Operator, role.Incomplete),
	AnnotateType("Parameter expansion operator '##'", nil, role.Operator, role.Incomplete),
	AnnotateType("Parameter expansion operator '%%'", nil, role.Operator, role.Incomplete),
	AnnotateType("Parameter expansion operator '::'", nil, role.Operator, role.Incomplete),
	AnnotateType("Parameter expansion operator ':'", nil, role.Operator, role.Incomplete),
	AnnotateType("Parameter expansion operator '#'", nil, role.Operator, role.Incomplete),
	AnnotateType("Parameter expansion operator '%'", nil, role.Operator, role.Incomplete),
	AnnotateType("Parameter expansion operator '//'", nil, role.Operator, role.Incomplete),
	AnnotateType("Parameter expansion operator '/'", nil, role.Operator, role.Incomplete),
	AnnotateType(">>", nil, role.Operator, role.Incomplete),
	AnnotateType("case pattern", nil, role.Expression, role.Case, role.Condition),
	AnnotateType("case pattern list", nil, role.Case, role.Body, role.Block),
	AnnotateType("case pattern", nil, role.Statement, role.Switch),
	AnnotateType("let command", nil, role.Expression, role.Assignment, role.Incomplete),
	AnnotateType("lazy LET expression", nil, role.Expression, role.Assignment, role.Incomplete),

	annotateTypeToken("cond_op", "-a", role.Operator, role.Relational, role.Equal),
	annotateTypeToken("cond_op", "-b", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-c", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-d", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-f", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-g", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-G", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-l", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-n", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-N", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-o", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-O", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-r", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-s", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-S", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-t", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-u", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-v", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-w", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-x", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-z", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-eq", role.Operator, role.Relational, role.Equal),
	annotateTypeToken("cond_op", "-ef", role.Operator, role.Relational, role.Equal, role.Incomplete),
	annotateTypeToken("cond_op", "-ne", role.Operator, role.Relational, role.Not, role.Equal),
	annotateTypeToken("cond_op", "-gt", role.Operator, role.Relational, role.GreaterThan),
	annotateTypeToken("cond_op", "-nt", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-ot", role.Operator, role.Relational, role.Incomplete),
	annotateTypeToken("cond_op", "-ge", role.Operator, role.Relational, role.GreaterThanOrEqual),
	annotateTypeToken("cond_op", "-lt", role.Operator, role.Relational, role.LessThan),
	annotateTypeToken("cond_op", "-le", role.Operator, role.Relational, role.LessThanOrEqual),
	annotateTypeToken("cond_op", "=", role.Operator, role.Relational, role.Equal),
	annotateTypeToken("cond_op", "<", role.Operator, role.Relational, role.LessThan),
	annotateTypeToken("cond_op", "<=", role.Operator, role.Relational, role.LessThanOrEqual),
	annotateTypeToken("cond_op", ">", role.Operator, role.Relational, role.GreaterThan),
	annotateTypeToken("cond_op", ">=", role.Operator, role.Relational, role.GreaterThanOrEqual),
	annotateTypeToken("cond_op", "!=", role.Operator, role.Relational, role.Not, role.Equal),
	annotateTypeToken("cond_op", "==", role.Operator, role.Relational, role.Equal),
	annotateTypeToken("cond_op", "=", role.Operator, role.Relational, role.Equal),

	AnnotateType("simple-command", nil, role.Expression),
	AnnotateType("var-def-element", nil, role.Expression, role.Declaration, role.Variable),
	AnnotateType("function-def-element", nil, role.Function, role.Declaration),
	AnnotateType("if shellcommand", nil, role.Statement, role.If),
	AnnotateType("for shellcommand", nil, role.Statement, role.For),
	AnnotateType("function", nil, role.Function, role.Declaration, role.Block),
	AnnotateType("named symbol", nil, role.Name, role.Identifier),
	AnnotateType("group element", nil, role.Body, role.Block),
	annotateTypeToken("generic bash command", "break", role.Statement, role.Break),
	annotateTypeToken("generic bash command", "continue", role.Statement, role.Continue),
	AnnotateType("include-command", nil, role.Statement, role.Import),
	AnnotateType("File reference", nil, role.Import, role.Pathname, role.Identifier),
	AnnotateType("=", nil, role.Operator, role.Assignment),

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

	AnnotateType("if shellcommand", MapObj(Obj{
		"children": Arr(ObjectRoles("condition")),
	}, Obj{
		"children": Arr(ObjectRoles("condition", role.If, role.Condition)),
	}), role.Statement, role.If),

	AnnotateType("elif", MapObj(Obj{
		"children": Arr(ObjectRoles("condition")),
	}, Obj{
		"children": Arr(ObjectRoles("condition", role.Else, role.If, role.Condition)),
	}), role.Statement, role.Else, role.If),

	AnnotateType("then", MapObj(Obj{
		"children": Arr(ObjectRoles("body")),
	}, Obj{
		"children": Arr(ObjectRoles("body", role.Then, role.Body, role.Block)),
	}), role.Then),

	AnnotateType("else", MapObj(Obj{
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
