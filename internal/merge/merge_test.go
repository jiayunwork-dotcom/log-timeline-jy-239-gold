package merge

import (
	"testing"
	"time"

	"log-timeline/internal/parse"
)

func TestMergeLengthAndOrder(t *testing.T) {
	base := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)
	a := []parse.Entry{{TS: base}, {TS: base.Add(2 * time.Minute)}}
	b := []parse.Entry{{TS: base.Add(1 * time.Minute)}, {TS: base.Add(3 * time.Minute)}}
	m := Merge(a, b)
	if len(m) != 4 {
		t.Fatalf("len=%d want 4", len(m))
	}
	for i := 1; i < len(m); i++ {
		if !m[i-1].TS.Before(m[i].TS) && !m[i-1].TS.Equal(m[i].TS) {
			t.Fatalf("not sorted at %d", i)
		}
	}
}

func TestSortStable(t *testing.T) {
	base := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)
	in := []parse.Entry{
		{TS: base, Message: "second"},
		{TS: base, Message: "first"},
	}
	out := Sort(in)
	if out[0].Message != "second" || out[1].Message != "first" {
		t.Fatalf("stable order broken: %+v", out)
	}
}

func TestSortDoesNotMutateInput(t *testing.T) {
	base := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)
	in := []parse.Entry{
		{TS: base.Add(time.Minute), Message: "later"},
		{TS: base, Message: "earlier"},
	}
	_ = Sort(in)
	if in[0].Message != "later" || in[1].Message != "earlier" {
		t.Fatalf("Sort mutated caller slice: %+v", in)
	}
}
