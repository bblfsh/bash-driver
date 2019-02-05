package fixtures

import (
	"path/filepath"
	"testing"

	"github.com/bblfsh/bash-driver/driver/normalizer"

	"gopkg.in/bblfsh/sdk.v2/driver"
	"gopkg.in/bblfsh/sdk.v2/driver/fixtures"
	"gopkg.in/bblfsh/sdk.v2/driver/native"
)

const projectRoot = "../../"

var Suite = &fixtures.Suite{
	Lang: "bash",
	Ext:  ".bash",
	Path: filepath.Join(projectRoot, fixtures.Dir),
	NewDriver: func() driver.Native {
		return native.NewDriverAt(filepath.Join(projectRoot, "build/bin/native"),
			native.UTF8)
	},
	Transforms: normalizer.Transforms,
	BenchName:  "very_long",
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
			"Comment",
			"file reference",
			"function-def-element",
		},
	},
	Docker: fixtures.DockerConfig{
		Image: "bash:latest",
	},
}

func TestBashDriver(t *testing.T) {
	Suite.RunTests(t)
}

func BenchmarkBashDriver(b *testing.B) {
	Suite.RunBenchmarks(b)
}
