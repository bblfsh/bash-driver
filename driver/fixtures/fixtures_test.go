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
	// XXX add function type
	Semantic: fixtures.SemanticConfig{
		BlacklistTypes: []string{
			"unevaluated string (STRING2)",
			"string",
			"string content",
			"backquote shellcommand",
			"File reference",
			"word",
			"variable",
			"assignment_word",
			"comment",
			"file reference",
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
