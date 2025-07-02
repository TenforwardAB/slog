package slog

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelNotice
	LevelWarn
	LevelError
	LevelCrit
	LevelAlert
	LevelEmerg
)

var levelNames = map[Level]string{
	LevelDebug:  "DEBUG",
	LevelInfo:   "INFO",
	LevelNotice: "NOTICE",
	LevelWarn:   "WARN",
	LevelError:  "ERROR",
	LevelCrit:   "CRIT",
	LevelAlert:  "ALERT",
	LevelEmerg:  "EMERG",
}

type logger struct {
	out      *log.Logger
	minLevel Level
}

var internal = &logger{
	out:      log.New(os.Stdout, "", log.LstdFlags),
	minLevel: LevelInfo,
}

func SetLevel(level interface{}) {
	switch v := level.(type) {
	case Level:
		internal.minLevel = v
	case string:
		parsed, err := parseLogLevel(v)
		if err != nil {
			panic(fmt.Sprintf("invalid log level: %s", v))
		}
		internal.minLevel = parsed
	default:
		panic(fmt.Sprintf("unsupported log level type: %T", level))
	}
}

func shouldLog(level Level) bool {
	return level >= internal.minLevel
}

func logMessage(level Level, msg string, args ...any) {
	if !shouldLog(level) {
		return
	}
	prefix := levelNames[level]
	internal.out.Printf("[%s] %s", prefix, fmt.Sprintf(msg, args...))
}

func parseLogLevel(level string) (Level, error) {
	switch strings.ToLower(level) {
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "notice":
		return LevelNotice, nil
	case "warn", "warning":
		return LevelWarn, nil
	case "error":
		return LevelError, nil
	case "crit", "critical":
		return LevelCrit, nil
	case "alert":
		return LevelAlert, nil
	case "emerg", "emergency":
		return LevelEmerg, nil
	default:
		return LevelInfo, fmt.Errorf("unknown log level: %s", level)
	}
}

func Emerg(msg string, args ...any)  { logMessage(LevelEmerg, msg, args...) }
func Alert(msg string, args ...any)  { logMessage(LevelAlert, msg, args...) }
func Crit(msg string, args ...any)   { logMessage(LevelCrit, msg, args...) }
func Error(msg string, args ...any)  { logMessage(LevelError, msg, args...) }
func Warn(msg string, args ...any)   { logMessage(LevelWarn, msg, args...) }
func Notice(msg string, args ...any) { logMessage(LevelNotice, msg, args...) }
func Info(msg string, args ...any)   { logMessage(LevelInfo, msg, args...) }
func Debug(msg string, args ...any)  { logMessage(LevelDebug, msg, args...) }
