package logger

import "github.com/rs/zerolog"

// Logger represents the logger interface
type Logger interface {
	Options() *Options

	Log() zerolog.Logger
}
