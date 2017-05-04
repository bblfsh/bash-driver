package normalizer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNativeToNoderVeryLong(t *testing.T) {
	require := require.New(t)

	f, err := getFixture("very_long.json")
	require.NoError(err)

	n, err := NativeToNoder.ToNode(f)
	require.NoError(err)
	require.NotNil(n)
}

const (
	fixtureDir = "fixtures"
)

func getFixture(name string) (data map[string]interface{}, err error) {
	path := filepath.Join(fixtureDir, name)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if errClose := f.Close(); err == nil {
			err = errClose
		}
	}()

	d := json.NewDecoder(f)
	if err := d.Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func TestNativeToNoderSheBang(t *testing.T) {
	require := require.New(t)

	f, err := getFixture("shebang.json")
	require.NoError(err)

	n, err := NativeToNoder.ToNode(f)
	require.NoError(err)
	require.NotNil(n)
	// do some checks
}

func TestNativeToNoderVar(t *testing.T) {
	require := require.New(t)

	f, err := getFixture("var.json")
	require.NoError(err)

	n, err := NativeToNoder.ToNode(f)
	require.NoError(err)
	require.NotNil(n)
	fmt.Println(n)
	// do some checks
}
