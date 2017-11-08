package normalizer

import (
	"github.com/bblfsh/bash-driver/driver/normalizer/intellij"

	"gopkg.in/bblfsh/sdk.v1/uast"
	. "gopkg.in/bblfsh/sdk.v1/uast/ann"
	"gopkg.in/bblfsh/sdk.v1/uast/transformer"
	"gopkg.in/bblfsh/sdk.v1/uast/transformer/annotatter"
	"gopkg.in/src-d/go-errors.v0"
)

// Transformers is the of list `transformer.Transfomer` to apply to a UAST, to
// learn more about the Transformers and the available ones take a look to:
// https://godoc.org/gopkg.in/bblfsh/sdk.v1/uast/transformers
var Transformers = []transformer.Tranformer{
	annotatter.NewAnnotatter(AnnotationRules),
}

var ErrRootMustBeFile = errors.NewKind("root must have internal type FILE")

// AnnotationRules describes how a UAST should be annotated with `uast.Role`.
//
// https://godoc.org/gopkg.in/bblfsh/sdk.v1/uast/ann
var AnnotationRules = On(Any).Self(
	On(Not(intellij.File)).Error(ErrRootMustBeFile.New()),
	On(intellij.File).Roles(uast.File).Descendants(
		On(intellij.Comment).Roles(uast.Comment),
		On(intellij.Shebang).Roles(uast.Comment, uast.Documentation),
		// variable declaration
		On(intellij.VarDefElement).Children(
			On(intellij.AssignmentWord).Roles(uast.Identifier),
		),
		// function declaration
		On(intellij.FunctionDefElement).Children(
			On(intellij.Function).Roles(uast.Function, uast.Declaration),
			On(intellij.NamedSymbol).Roles(uast.Function, uast.Declaration, uast.Name),
			On(intellij.GroupElement).Roles(uast.Function, uast.Declaration, uast.Body, uast.Block),
		),
		// if statement
		On(intellij.IfShellCommand).Roles(uast.If, uast.Statement),
		// for statement
		On(intellij.ForShellCommand).Roles(uast.For, uast.Statement),
		// while and until statement
		On(intellij.WhileLoop).Roles(uast.While, uast.Statement),
		On(intellij.UntilLoop).Roles(uast.While, uast.Statement),
	),
)
