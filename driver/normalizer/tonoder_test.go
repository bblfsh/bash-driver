package normalizer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToNoderVeryLong(t *testing.T) {
	require := require.New(t)

	f, err := getFixture(integration, "very_long.native")
	require.NoError(err)

	n, err := NativeToNoder.ToNode(f)
	require.NoError(err)
	require.NotNil(n)
}
