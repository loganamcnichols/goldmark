package extension

import (
	"testing"

	"github.com/loganamcnichols/goldmark"
	"github.com/loganamcnichols/goldmark/renderer/html"
	"github.com/loganamcnichols/goldmark/testutil"
)

func TestDefinitionList(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			DefinitionList,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/definition_list.txt", t, testutil.ParseCliCaseArg()...)
}
