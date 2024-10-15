package extension

import (
	"testing"

	"github.com/loganamcnichols/goldmark"
	"github.com/loganamcnichols/goldmark/renderer/html"
	"github.com/loganamcnichols/goldmark/testutil"
)

func TestTaskList(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			TaskList,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/tasklist.txt", t, testutil.ParseCliCaseArg()...)
}
