package parse

import (
	"fmt"
	"strings"
	"time"
)

type Entry struct {
	TS      time.Time
	Level   string
	Source  string
	Message string
}

func ParseLine(line string) (Entry, error) {
	line = strings.TrimRight(line, "\r\n")
	line = strings.TrimSpace(line)
	if line == "" {
		return Entry{}, fmt.Errorf("parse: empty line")
	}
	fields := strings.SplitN(line, " ", 3)
	if len(fields) < 2 {
		return Entry{}, fmt.Errorf("parse: need timestamp and level")
	}
	ts, err := time.Parse(time.RFC3339, fields[0])
	if err != nil {
		return Entry{}, fmt.Errorf("parse: bad timestamp %q: %w", fields[0], err)
	}
	level := strings.ToUpper(fields[1])
	rest := ""
	if len(fields) == 3 {
		rest = fields[2]
	}
	src := ""
	msg := rest
	if strings.HasPrefix(rest, "[") {
		if end := strings.Index(rest, "]"); end > 0 {
			src = rest[1:end]
			msg = strings.TrimSpace(rest[end+1:])
		}
	}
	return Entry{TS: ts, Level: level, Source: src, Message: msg}, nil
}
