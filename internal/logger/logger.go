package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type logger struct {
	opt *Options
	log *zerolog.Logger
}

// New creates a new logger instance
func New(opts ...Option) (Logger, error) {
	opt, err := newOptions(opts...)
	if err != nil {
		return nil, err
	}

	// Init logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	lgr := log.With().Logger()

	// Set format
	lgr = convertFormat(lgr, opt)

	// Set level
	level, _ := zerolog.ParseLevel(opt.Level)
	zerolog.SetGlobalLevel(level)
	if opt.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	return &logger{
		opt: opt,
		log: &lgr,
	}, nil
}

// Options returns the list of options
func (l *logger) Options() *Options {
	return l.opt
}

// Log returns the zerolog logger
func (l *logger) Log() zerolog.Logger {
	return *l.log
}

func convertFormat(lgr zerolog.Logger, opt *Options) zerolog.Logger {
	return lgr.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	})
}
