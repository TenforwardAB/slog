//go:build dev

package slog

func Crazy(msg string, args ...any) {
	logMessage(LevelDebug, "[CRAZY] "+msg, args...)
}
