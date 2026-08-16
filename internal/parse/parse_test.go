package parse

import (
	"testing"
)

func TestParseLineOK(t *testing.T) {
	e, err := ParseLine("2026-01-01T10:00:00Z INFO [api] request received")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.Level != "INFO" || e.Source != "api" || e.Message != "request received" {
		t.Fatalf("parsed wrong: %+v", e)
	}
}

func TestParseLineBadTimestamp(t *testing.T) {
	if _, err := ParseLine("not-a-time INFO msg"); err == nil {
		t.Fatal("expected error for bad timestamp")
	}
}

func TestParseLineEmpty(t *testing.T) {
	if _, err := ParseLine("   "); err == nil {
		t.Fatal("expected error for empty line")
	}
}

func TestParseLineNoSource(t *testing.T) {
	e, err := ParseLine("2026-01-01T10:00:00Z ERROR something broke")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.Source != "" || e.Message != "something broke" {
		t.Fatalf("parsed wrong: %+v", e)
	}
}

func TestParseLineRequiresLevel(t *testing.T) {
	if _, err := ParseLine("2026-01-01T10:00:00Z"); err == nil {
		t.Fatal("expected error when level is missing")
	}
}

func TestParseLineNormalizesLevel(t *testing.T) {
	e, err := ParseLine("2026-01-01T10:00:00Z info [api] hi")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.Level != "INFO" {
		t.Fatalf("Level=%q want INFO", e.Level)
	}
}
