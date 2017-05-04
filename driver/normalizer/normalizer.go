package normalizer

import (
	"github.com/bblfsh/sdk/uast"
	. "github.com/bblfsh/sdk/uast/ann"
)

var AnnotationRules *Rule = nil

var NativeToNoder = &uast.BaseToNoder{
	InternalTypeKey:    "type",
	OffsetKey:          "startOffset",
	TopLevelIsRootNode: true,
	TokenKeys:          map[string]bool{"text": true},
}
