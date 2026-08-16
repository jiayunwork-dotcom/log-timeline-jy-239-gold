package merge

import (
	"sort"

	"log-timeline/internal/parse"
)

func Sort(entries []parse.Entry) []parse.Entry {
	out := append([]parse.Entry(nil), entries...)
	sort.SliceStable(out, func(i, j int) bool {
		return out[i].TS.Before(out[j].TS)
	})
	return out
}

func Merge(a, b []parse.Entry) []parse.Entry {
	merged := append([]parse.Entry(nil), a...)
	merged = append(merged, b...)
	return Sort(merged)
}
