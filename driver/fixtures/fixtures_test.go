package fixtures

import (
	"path/filepath"
	"testing"

	"github.com/bblfsh/bash-driver/driver/normalizer"
	"gopkg.in/bblfsh/sdk.v2/sdk/driver"
	"gopkg.in/bblfsh/sdk.v2/sdk/driver/fixtures"
)

const projectRoot = "../../"

var Suite = &fixtures.Suite{
	Lang: "bash",
	Ext:  ".bash",
	Path: filepath.Join(projectRoot, fixtures.Dir),
	NewDriver: func() driver.BaseDriver {
		return driver.NewExecDriverAt(filepath.Join(projectRoot, "build/bin/native"))
	},
	Transforms: normalizer.Transforms,
	BenchName: "very_long",
	Semantic: fixtures.SemanticConfig{
		BlacklistTypes: []string{
			// TODO: list native types that should be converted to semantic UAST
		},
	},
	// TODO XXX: not working ("fixtures.test" not found)
	//Docker:fixtures.DockerConfig{
		//Image:"bash:latest",
	//},
}

func TestBashDriver(t *testing.T) {
	Suite.RunTests(t)
}

func BenchmarkBashDriver(b *testing.B) {
	Suite.RunBenchmarks(b)
}
