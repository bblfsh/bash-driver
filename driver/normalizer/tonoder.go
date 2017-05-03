package normalizer

import (
	"github.com/bblfsh/sdk/uast"
)

var NativeToNoder = &uast.BaseToNoder{
	InternalTypeKey:    "elementType",
	OffsetKey:          "startOffset",
	TopLevelIsRootNode: true,
	TokenKeys:          map[string]bool{"text": true},
}
