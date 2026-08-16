package timeline

import (
	"testing"
	"time"

	"log-timeline/internal/parse"
)

func mk(ts time.Time, src string) parse.Entry {
	return parse.Entry{TS: ts, Source: src}
}

func TestBuildGapThreshold(t *testing.T) {
	base := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)
	entries := []parse.Entry{
		mk(base, "a"),
		mk(base.Add(4*time.Minute), "a"),
		mk(base.Add(10*time.Minute), "a"),
	}
	tl := Build(entries, 5*time.Minute)
	if len(tl.Gaps) != 1 {
		t.Fatalf("gaps=%d want 1 (the 6m gap; 4m gap below threshold)", len(tl.Gaps))
	}
}

func TestBuildGapBoundary(t *testing.T) {
	base := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)
	entries := []parse.Entry{
		mk(base, "a"),
		mk(base.Add(5*time.Minute), "a"),
	}
	tl := Build(entries, 5*time.Minute)
	if len(tl.Gaps) != 1 {
		t.Fatalf("gap exactly == threshold should be flagged (>=), got %d", len(tl.Gaps))
	}
}

func TestBySourceUnknown(t *testing.T) {
	base := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)
	tl := Build([]parse.Entry{mk(base, "api")}, time.Minute)
	got := BySource(tl, "nonexistent")
	if got == nil {
		t.Fatal("BySource should return non-nil empty slice for unknown source")
	}
	if len(got) != 0 {
		t.Fatalf("want 0 entries, got %d", len(got))
	}
}

func TestBuildSortsBeforeGaps(t *testing.T) {
	base := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)
	entries := []parse.Entry{
		mk(base.Add(10*time.Minute), "a"),
		mk(base, "a"),
		mk(base.Add(4*time.Minute), "a"),
	}
	tl := Build(entries, 5*time.Minute)
	if len(tl.Entries) != 3 || !tl.Entries[0].TS.Equal(base) {
		t.Fatalf("entries not time-sorted: %+v", tl.Entries)
	}
	if len(tl.Gaps) != 1 {
		t.Fatalf("gaps=%d want 1 after sorting", len(tl.Gaps))
	}
}

func TestBuildGapReportsSeconds(t *testing.T) {
	base := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)
	entries := []parse.Entry{
		mk(base, "a"),
		mk(base.Add(90*time.Second), "a"),
	}
	tl := Build(entries, time.Minute)
	if len(tl.Gaps) != 1 {
		t.Fatalf("gaps=%d want 1", len(tl.Gaps))
	}
	if tl.Gaps[0].Seconds != 90 {
		t.Fatalf("Seconds=%v want 90", tl.Gaps[0].Seconds)
	}
}
