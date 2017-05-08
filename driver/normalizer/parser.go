package normalizer

import (
	"github.com/bblfsh/sdk/protocol/driver"
	"github.com/bblfsh/sdk/protocol/native"
)

var ToNoder = &native.ObjectToNoder{
	InternalTypeKey:    "elementType",
	OffsetKey:          "startOffset",
	TopLevelIsRootNode: true,
	TokenKeys:          map[string]bool{"text": true},
}

// ASTParserBuilder creates a parser that transform source code files
// into *uast.Node.
func ASTParserBuilder(opts driver.ASTParserOptions) (
	driver.ASTParser, error) {

	parser, err := native.ExecParser(ToNoder, opts.NativeBin)
	if err != nil {
		return nil, err
	}

	return parser, nil
}
