package main

import "github.com/TenforwardAB/slog"

func main() {
	slog.SetLevel("info") // or "debug", "warn", etc.

	slog.Debug("Debugging enabled")
	slog.Info("Information message")
	slog.Notice("Notice: something noteworthy")
	slog.Warn("Warning: something might go wrong")
	slog.Error("Error occurred")
	slog.Crit("Critical failure")
	slog.Alert("Immediate attention required")
	slog.Emerg("System unusable")

	slog.Crazy("Dev-only logging")
}
