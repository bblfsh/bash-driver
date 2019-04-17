package fixtures

import (
	"path/filepath"
	"testing"

	"github.com/bblfsh/bash-driver/driver/normalizer"

	"github.com/bblfsh/sdk/v3/driver"
	"github.com/bblfsh/sdk/v3/driver/fixtures"
	"github.com/bblfsh/sdk/v3/driver/native"
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
			"unevaluated_string2",
			"string_content",
			"File_reference",
			"word",
			"variable",
			"assignment_word",
			"Comment",
			"file_reference",
			"function-def-element",
		},
	},
}

func TestBashDriver(t *testing.T) {
	Suite.RunTests(t)
}

func BenchmarkBashDriver(b *testing.B) {
	Suite.RunBenchmarks(b)
}
