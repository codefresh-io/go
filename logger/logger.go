package logger

import (
	"io"
	"os"

	"github.com/inconshreveable/log15"
)

type (
	// Logger exposes functions to log messeges
	Logger interface {
		log15.Logger
	}

	// Options to create new logger
	Options struct {
		DisableStdout    bool
		Context          []interface{}
		File             io.Writer
		StdoutJSONFormat bool
	}
)

// New build logger
func New(options Options) (Logger, error) {
	l := log15.New(options.Context...)
	if options.File != nil {
		l.SetHandler(log15.StreamHandler(options.File, log15.JsonFormat()))
	}
	if options.DisableStdout {
		return l, nil
	}

	if options.StdoutJSONFormat {
		l.SetHandler(log15.StreamHandler(os.Stdout, log15.JsonFormat()))
	}
	return l, nil
}
