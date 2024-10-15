package extension

import (
	"testing"

	"github.com/loganamcnichols/goldmark"
	"github.com/loganamcnichols/goldmark/renderer/html"
	"github.com/loganamcnichols/goldmark/testutil"
)

func TestTypographer(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			Typographer,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/typographer.txt", t, testutil.ParseCliCaseArg()...)
}
