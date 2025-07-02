//go:build dev

package slog

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestCrazyLog(t *testing.T) {
	var buf bytes.Buffer
	original := internal.out
	internal.out = log.New(&buf, "", 0)
	defer func() { internal.out = original }()

	SetLevel(LevelDebug)
	Crazy("this is a crazy test")

	out := buf.String()
	if !strings.Contains(out, "[CRAZY] this is a crazy test") {
		t.Errorf("Crazy log not found or incorrect: %s", out)
	}
}
