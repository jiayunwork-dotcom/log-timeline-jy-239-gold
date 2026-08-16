package report

import (
	"bufio"
	"fmt"
	"io"
	"time"

	"log-timeline/internal/timeline"
)

func Write(w io.Writer, tl timeline.Timeline) (err error) {
	bw := bufio.NewWriter(w)
	defer func() {
		if ferr := bw.Flush(); ferr != nil && err == nil {
			err = ferr
		}
	}()
	fmt.Fprintf(bw, "entries: %d\n", len(tl.Entries))
	fmt.Fprintf(bw, "gaps: %d\n", len(tl.Gaps))
	for i, e := range tl.Entries {
		fmt.Fprintf(bw, "%d %s %s %s %s\n", i, e.TS.Format(time.RFC3339), e.Level, e.Source, e.Message)
	}
	for _, g := range tl.Gaps {
		fmt.Fprintf(bw, "GAP %.0fs %s -> %s\n", g.Seconds, g.From.Format(time.RFC3339), g.To.Format(time.RFC3339))
	}
	return nil
}
