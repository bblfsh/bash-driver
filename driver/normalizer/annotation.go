package normalizer

import (
	. "github.com/bblfsh/sdk/uast"
	. "github.com/bblfsh/sdk/uast/ann"

	"github.com/bblfsh/bash-driver/driver/normalizer/intellij"
	"srcd.works/go-errors.v0"
)

var (
	ErrRootMustBeFile = errors.NewKind("root must have internal type FILE")
)

var AnnotationRules = On(Any).Self(
	On(Not(intellij.File)).Error(ErrRootMustBeFile.New()),
	On(intellij.File).Roles(File).Descendants(
		On(intellij.Comment).Roles(Comment),
		On(intellij.Shebang).Roles(Comment, Documentation),
		// variable declaration
		On(intellij.VarDefElement).Children(
			On(intellij.AssignmentWord).Roles(SimpleIdentifier),
		),
		// function declaration
		On(intellij.FunctionDefElement).Children(
			On(intellij.Function).Roles(FunctionDeclaration),
			On(intellij.NamedSymbol).Roles(FunctionDeclarationName),
			On(intellij.GroupElement).Roles(FunctionDeclarationBody, Block),
		),
		// if statement
		On(intellij.IfShellCommand).Roles(If, Statement),
		// for statement
		On(intellij.ForShellCommand).Roles(ForEach, Statement),
		// while and until statement
		On(intellij.WhileLoop).Roles(While, Statement),
		On(intellij.UntilLoop).Roles(While, Statement),
	),
)
