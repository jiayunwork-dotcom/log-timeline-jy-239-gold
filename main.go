package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"log-timeline/internal/parse"
	"log-timeline/internal/report"
	"log-timeline/internal/timeline"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "log-timeline:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fs := flag.NewFlagSet("log-timeline", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	input := fs.String("input", "-", "log file ('-' for stdin)")
	gap := fs.String("gap", "5m", "gap threshold duration (e.g. 5m, 30s)")
	out := fs.String("out", "-", "write timeline ('-' for stdout)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	thr, err := time.ParseDuration(*gap)
	if err != nil {
		return fmt.Errorf("parse gap: %w", err)
	}
	var r io.Reader = os.Stdin
	if *input != "-" {
		f, err := os.Open(*input)
		if err != nil {
			return fmt.Errorf("open input: %w", err)
		}
		defer f.Close()
		r = f
	}
	var entries []parse.Entry
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		e, err := parse.ParseLine(line)
		if err != nil {
			return fmt.Errorf("parse line: %w", err)
		}
		entries = append(entries, e)
	}
	if err := sc.Err(); err != nil {
		return fmt.Errorf("read input: %w", err)
	}
	tl := timeline.Build(entries, thr)
	var w io.Writer = os.Stdout
	if *out != "-" {
		of, err := os.Create(*out)
		if err != nil {
			return fmt.Errorf("create out: %w", err)
		}
		defer of.Close()
		w = of
	}
	if err := report.Write(w, tl); err != nil {
		return fmt.Errorf("write report: %w", err)
	}
	return nil
}
