package report

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"time"

	"log-timeline/internal/parse"
	"log-timeline/internal/timeline"
)

type failWriter struct{ fail bool }

func (f failWriter) Write(p []byte) (int, error) {
	if f.fail {
		return 0, errors.New("write fail")
	}
	return len(p), nil
}

func TestWriteOK(t *testing.T) {
	base := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)
	tl := timeline.Timeline{Entries: []parse.Entry{{TS: base, Level: "INFO", Source: "api", Message: "hi"}}}
	var buf bytes.Buffer
	if err := Write(&buf, tl); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), "entries: 1") {
		t.Fatalf("missing content: %q", buf.String())
	}
}

func TestWriteFlushError(t *testing.T) {
	tl := timeline.Timeline{Entries: []parse.Entry{{TS: time.Now()}}}
	if err := Write(failWriter{true}, tl); err == nil {
		t.Fatal("expected flush error to propagate")
	}
}
