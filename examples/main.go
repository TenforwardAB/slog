package main

import "github.com/TenforwardAB/slog"

func main() {
	slog.SetLevel(slog.LevelDebug)

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
