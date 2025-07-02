//go:build !dev

package slog

func Crazy(msg string, args ...any) {
	// no-op in production
}
