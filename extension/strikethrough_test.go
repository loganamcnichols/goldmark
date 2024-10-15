package extension

import (
	"testing"

	"github.com/loganamcnichols/goldmark"
	"github.com/loganamcnichols/goldmark/renderer/html"
	"github.com/loganamcnichols/goldmark/testutil"
)

func TestStrikethrough(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			Strikethrough,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/strikethrough.txt", t, testutil.ParseCliCaseArg()...)
}
