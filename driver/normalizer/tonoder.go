package normalizer

import (
	"github.com/bblfsh/sdk/uast"
)

var NativeToNoder = &uast.BaseToNoder{
	InternalTypeKey:    "type",
	OffsetKey:          "startOffset",
	TopLevelIsRootNode: true,
}
