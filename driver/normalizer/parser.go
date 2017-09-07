package normalizer

import (
	"gopkg.in/bblfsh/sdk.v0/protocol/driver"
	"gopkg.in/bblfsh/sdk.v0/protocol/native"
)

var ToNoder = &native.ObjectToNoder{
	InternalTypeKey:    "elementType",
	OffsetKey:          "startOffset",
	TopLevelIsRootNode: true,
	TokenKeys:          map[string]bool{"text": true},
}

// ParserBuilder creates a parser that transform source code files into *uast.Node.
func ParserBuilder(opts driver.ParserOptions) (driver.Parser, error) {
	parser, err := native.ExecParser(ToNoder, opts.NativeBin)
	if err != nil {
		return nil, err
	}

	return parser, nil
}
