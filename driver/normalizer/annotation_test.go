package normalizer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/bblfsh/sdk/protocol/native"
	"github.com/bblfsh/sdk/uast"
	"github.com/bblfsh/sdk/uast/ann"
	"github.com/stretchr/testify/require"
)

// the directories with the fixtures for the integration and the unit
// tests (we will be reusing some fixtures from the integration tests
// in this unit tests).
const (
	integration = "../../tests"
	unit        = "fixtures"
)

// Reads a native AST encoded in JSON from a file in the fixture directory.
func getFixture(dir, file string) (data map[string]interface{}, err error) {
	path := filepath.Join(dir, file)
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

	ast, ok := data["ast"]
	if !ok {
		return nil, fmt.Errorf("ast object not found")
	}

	asMap, ok := ast.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("cannot convert ast to map")
	}

	return asMap, nil
}

// Reads a native AST encoded in JSON from a file in the fixture directory, runs
// NativeToNoder on it and annotate it with AnnotationsRules.
func annotateFixture(dir, file string) (*uast.Node, error) {
	return annotateFixtureWith(dir, file, ToNoder, AnnotationRules)
}

// The same as annotateFixture above but using the ToNoder and the
// Annotation rules provided as argguments.
func annotateFixtureWith(
	dir, file string, toNoder *native.ObjectToNoder, rules *ann.Rule) (
	*uast.Node, error) {

	f, err := getFixture(dir, file)
	if err != nil {
		return nil, err
	}

	n, err := toNoder.ToNode(f)
	if err != nil {
		return nil, err
	}

	err = rules.Apply(n)
	if err != nil {
		return nil, err
	}

	return n, err
}

// return an slice with all the nodes in the tree that contains the role
// at least once.
func find(tree *uast.Node, role uast.Role) []*uast.Node {
	var found []*uast.Node
	_find(tree, role, &found)
	return found
}

func _find(n *uast.Node, r uast.Role, ret *[]*uast.Node) {
	for _, e := range n.Roles {
		if e == r {
			*ret = append(*ret, n)
			break
		}
	}
	for _, child := range n.Children {
		_find(child, r, ret)
	}
}

func mustBeTheSame(t *testing.T, expected, obtained []string) {
	sort.Strings(expected)
	sort.Strings(obtained)
	require.Equal(t, expected, obtained)
}

// returns a slice whith the tokens in the given nodes.
func tokens(s ...*uast.Node) []string {
	var ret []string
	for _, e := range s {
		ret = append(ret, e.Token)
	}
	return ret
}

func TestAnnotationsErrorIfRootIsNotFile(t *testing.T) {
	require := require.New(t)
	_, err := annotateFixture(unit, "root_is_not_file.json")
	require.Error(err)

	detailedError, ok := err.(ann.RuleError)
	require.True(ok)
	require.True(ErrRootMustBeFile.Is(detailedError.Inner()))
}

func TestAnnotationsRootIsFile(t *testing.T) {
	require := require.New(t)
	n, err := annotateFixture(integration, "var_declaration.native")
	require.NoError(err)
	require.Contains(n.Roles, uast.File)
}

func TestAnnotationsCommentAreComments(t *testing.T) {
	n, err := annotateFixture(integration, "comments.native")
	require.NoError(t, err)

	expected := []string{
		"# comment 1",
		"# comment 2",
	}
	obtained := tokens(find(n, uast.Comment)...)
	mustBeTheSame(t, expected, obtained)
}
func TestAnnotationsShebangIsComment(t *testing.T) {
	n, err := annotateFixture(integration, "shebang.native")
	require.NoError(t, err)

	expected := []string{"#!/bin/bash\n"}
	obtained := tokens(find(n, uast.Comment)...)
	mustBeTheSame(t, expected, obtained)
}

func TestAnnotationsShebangIsDocumentation(t *testing.T) {
	n, err := annotateFixture(integration, "shebang.native")
	require.NoError(t, err)

	expected := []string{"#!/bin/bash\n"}
	obtained := tokens(find(n, uast.Documentation)...)
	mustBeTheSame(t, expected, obtained)
}

func TestAnnotationsOrdinatyCommentsAreNotDocumentation(t *testing.T) {
	n, err := annotateFixture(integration, "comments.native")
	require.NoError(t, err)

	var expected []string // we don't expect to find any documentation in the file
	obtained := tokens(find(n, uast.Documentation)...)
	mustBeTheSame(t, expected, obtained)
}

func TestAnnotationsVariableDeclaration(t *testing.T) {
	n, err := annotateFixture(integration, "var_declaration.native")
	require.NoError(t, err)

	var expected = []string{"a"}
	obtained := tokens(find(n, uast.SimpleIdentifier)...)
	mustBeTheSame(t, expected, obtained)
}

func TestAnnotationsFunctionDeclaration(t *testing.T) {
	n, err := annotateFixture(integration, "function_declaration.native")
	require.NoError(t, err)

	var expected = []string{"function"}
	obtained := tokens(find(n, uast.FunctionDeclaration)...)
	mustBeTheSame(t, expected, obtained)

	expected = []string{"foo"}
	obtained = tokens(find(n, uast.FunctionDeclarationName)...)
	mustBeTheSame(t, expected, obtained)

	bodies := find(n, uast.FunctionDeclarationBody)
	require.Equal(t, 1, len(bodies))

	blocks := find(n, uast.Block)
	require.Equal(t, 1, len(blocks))
}
