package timeline

import (
	"sort"
	"time"

	"log-timeline/internal/parse"
)

type Gap struct {
	From    time.Time
	To      time.Time
	Seconds float64
}

type Timeline struct {
	Entries []parse.Entry
	Gaps    []Gap
}

func Build(entries []parse.Entry, gapThreshold time.Duration) Timeline {
	sorted := append([]parse.Entry(nil), entries...)
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].TS.Before(sorted[j].TS)
	})
	tl := Timeline{Entries: sorted}
	if len(sorted) < 2 {
		return tl
	}
	for i := 1; i < len(sorted); i++ {
		d := sorted[i].TS.Sub(sorted[i-1].TS)
		if d >= gapThreshold {
			tl.Gaps = append(tl.Gaps, Gap{
				From:    sorted[i-1].TS,
				To:      sorted[i].TS,
				Seconds: d.Seconds(),
			})
		}
	}
	return tl
}

func BySource(tl Timeline, src string) []parse.Entry {
	out := make([]parse.Entry, 0)
	for _, e := range tl.Entries {
		if e.Source == src {
			out = append(out, e)
		}
	}
	return out
}
