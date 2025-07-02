package slog

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	original := internal.out
	internal.out = log.New(&buf, "", 0)
	f()
	internal.out = original
	return buf.String()
}

func TestParseLogLevel(t *testing.T) {
	cases := map[string]Level{
		"debug":     LevelDebug,
		"info":      LevelInfo,
		"notice":    LevelNotice,
		"warn":      LevelWarn,
		"warning":   LevelWarn,
		"error":     LevelError,
		"crit":      LevelCrit,
		"critical":  LevelCrit,
		"alert":     LevelAlert,
		"emerg":     LevelEmerg,
		"emergency": LevelEmerg,
	}

	for input, expected := range cases {
		got, err := parseLogLevel(input)
		if err != nil {
			t.Errorf("unexpected error for %s: %v", input, err)
		}
		if got != expected {
			t.Errorf("expected %v for %s, got %v", expected, input, got)
		}
	}

	_, err := parseLogLevel("banana")
	if err == nil {
		t.Errorf("expected error for unknown level")
	}
}

func TestSetLevelFromString(t *testing.T) {
	SetLevel("debug")
	if internal.minLevel != LevelDebug {
		t.Errorf("SetLevel failed to set debug level")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for invalid log level")
		}
	}()
	SetLevel("invalid")
}

func TestSetLevelFromEnum(t *testing.T) {
	SetLevel(LevelNotice)
	if internal.minLevel != LevelNotice {
		t.Errorf("SetLevel failed to set notice level")
	}
}

func TestSetLevelFromUnsupportedType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for unsupported type")
		}
	}()
	SetLevel(123)
}

func TestLogFiltering(t *testing.T) {
	SetLevel(LevelError)

	out := captureOutput(func() {
		Debug("should not be logged")
		Error("should be logged")
		Emerg("should also be logged")
	})

	if strings.Contains(out, "DEBUG") {
		t.Errorf("Debug should not be logged")
	}
	if !strings.Contains(out, "ERROR") || !strings.Contains(out, "EMERG") {
		t.Errorf("Expected logs missing: %s", out)
	}
}

func TestAllLogFunctions(t *testing.T) {
	SetLevel(LevelDebug)

	tests := []struct {
		name     string
		fn       func(string, ...any)
		expected string
	}{
		{"Debug", Debug, "DEBUG"},
		{"Info", Info, "INFO"},
		{"Notice", Notice, "NOTICE"},
		{"Warn", Warn, "WARN"},
		{"Error", Error, "ERROR"},
		{"Crit", Crit, "CRIT"},
		{"Alert", Alert, "ALERT"},
		{"Emerg", Emerg, "EMERG"},
	}

	for _, tt := range tests {
		out := captureOutput(func() {
			tt.fn("test %s", tt.name)
		})
		if !strings.Contains(out, tt.expected) {
			t.Errorf("%s log not found in output: %s", tt.expected, out)
		}
	}
}
