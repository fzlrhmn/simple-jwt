package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Level indicates logger level
type Level int8

// Minimum required level enum
const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

// TranslateLevel translates level in string to Level type
// If none matches, default to WarnLevel and return not ok
func TranslateLevel(levelStr string) (level Level, ok bool) {
	switch levelStr {
	case "debug":
		return DebugLevel, true
	case "info":
		return InfoLevel, true
	case "warn":
		return WarnLevel, true
	case "error":
		return ErrorLevel, true
	}

	return ErrorLevel, false
}

// GenerateConfig for generate zap config for logging
func GenerateConfig(env string, level Level) *zap.Config {
	var cfg zap.Config

	// Initial config depends on env
	if env == "development" {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.EncoderConfig.EncodeTime = nil
	} else {
		cfg = zap.NewProductionConfig()
	}

	// Set level if not empty
	switch level {
	case DebugLevel:
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case InfoLevel:
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case WarnLevel:
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case ErrorLevel:
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	return &cfg
}
