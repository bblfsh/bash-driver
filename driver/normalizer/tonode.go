package normalizer

import "gopkg.in/bblfsh/sdk.v1/uast"

// ToNode is an instance of `uast.ObjectToNode`, defining how to transform an
// into a UAST (`uast.Node`).
//
// https://godoc.org/gopkg.in/bblfsh/sdk.v1/uast#ObjectToNode
var ToNode = &uast.ObjectToNode{
	InternalTypeKey:    "elementType",
	OffsetKey:          "startOffset",
	EndOffsetKey:       "endOffset",
	TopLevelIsRootNode: true,
	TokenKeys:          map[string]bool{"text": true},
	Modifier: func(n map[string]interface{}) error {
		// Create endOffset as startOffset + textLength
		if textEndPos, ok := n["textLength"].(float64); ok {
			n["endOffset"] = textEndPos + n["startOffset"].(float64)
			delete(n, "textLength")
		}

		return nil
	},
}
