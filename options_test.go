package goldmark_test

import (
	"testing"

	. "github.com/loganamcnichols/goldmark"
	"github.com/loganamcnichols/goldmark/parser"
	"github.com/loganamcnichols/goldmark/testutil"
)

func TestAttributeAndAutoHeadingID(t *testing.T) {
	markdown := New(
		WithParserOptions(
			parser.WithAttribute(),
			parser.WithAutoHeadingID(),
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/options.txt", t, testutil.ParseCliCaseArg()...)
}
