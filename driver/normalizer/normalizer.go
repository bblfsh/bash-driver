package normalizer

import (
	. "github.com/bblfsh/sdk/uast"
	. "github.com/bblfsh/sdk/uast/ann"
)

const (
	ErrRootMustBeFile = "root must have internal type FILE"
)

var AnnotationRules *Rule = On(Any).Self(
	On(Not(HasInternalType("FILE"))).Error(ErrRootMustBeFile),
	On(HasInternalType("FILE")).Roles(File).Descendants(
		On(HasInternalType("[Bash] Comment")).Roles(Comment),
		On(HasInternalType("[Bash] shebang element")).Roles(Comment, Documentation),
		// variable declaration
		On(HasInternalType("var-def-element")).Children(
			On(HasInternalType("[Bash] assignment_word")).Roles(SimpleIdentifier)),
		// function declaration
		On(HasInternalType("function-def-element")).Children(
			On(HasInternalType("[Bash] function")).Roles(FunctionDeclaration),
			On(HasInternalType("[Bash] named symbol")).Roles(FunctionDeclarationName),
			On(HasInternalType("group element")).Roles(FunctionDeclarationBody, Block)),
	),
)
