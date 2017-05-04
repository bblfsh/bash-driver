package normalizer

import (
	"sort"
	"testing"

	"github.com/bblfsh/sdk/uast"
	"github.com/bblfsh/sdk/uast/ann"
	"github.com/stretchr/testify/require"
)

// Reads the contents of a file form the fixture directory, runs
// NativeToNoder on it and annotate it with AnnotationsRules.
func annotateFixture(path string) (*uast.Node, error) {
	return annotateFixtureWith(path, NativeToNoder, AnnotationRules)
}

// Reads the contents of a file form the fixture directory, runs
// the given ToNoder on it and annotate it with the given annotation rules.
func annotateFixtureWith(
	path string, toNoder *uast.BaseToNoder, rules *ann.Rule) (
	*uast.Node, error) {

	f, err := getFixture(path)
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

func TestAnnotationsErrorIfRootIsNotFile(t *testing.T) {
	require := require.New(t)
	_, err := annotateFixture("root_is_not_file.json")
	require.Error(err)
	require.EqualError(err, ErrRootMustBeFile)
}

func TestAnnotationsRootIsFile(t *testing.T) {
	require := require.New(t)
	n, err := annotateFixture("var.json")
	require.NoError(err)
	require.Contains(n.Roles, uast.File)
}

func TestAnnotationsCommentAreComments(t *testing.T) {
	n, err := annotateFixture("comments.json")
	require.NoError(t, err)

	expected := []string{
		"# comment 1",
		"# comment 2",
	}
	comments := tokens(find(n, uast.Comment)...)
	require.Equal(t, expected, comments)
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

// returns a slice whith the tokens in the given nodes, sorted alphabetically.
func tokens(s ...*uast.Node) []string {
	var ret []string
	for _, e := range s {
		ret = append(ret, e.Token)
	}
	sort.Strings(ret)
	return ret
}

func TestAnnotationsShebangIsComment(t *testing.T) {
	n, err := annotateFixture("shebang.json")
	require.NoError(t, err)

	expected := []string{"#!/bin/bash\n"}
	comments := tokens(find(n, uast.Comment)...)
	require.Equal(t, expected, comments)
}

func TestAnnotationsShebangIsDocumentation(t *testing.T) {
	n, err := annotateFixture("shebang.json")
	require.NoError(t, err)

	expected := []string{"#!/bin/bash\n"}
	documentation := tokens(find(n, uast.Documentation)...)
	require.Equal(t, expected, documentation)
}

func TestAnnotationsOrdinatyCommentsAreNotDocumentation(t *testing.T) {
	n, err := annotateFixture("comments.json")
	require.NoError(t, err)

	var expected []string // we don't expect to find any documentation in the file
	documentation := tokens(find(n, uast.Documentation)...)
	require.Equal(t, expected, documentation)
}
