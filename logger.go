package main

import (
	"github.com/op/go-logging"
	"os"
)

// CreateLogger creates a logger and sets the backend formatter for logging.
// It returns the logger.
func CreateLogger() *logging.Logger {
	log := logging.MustGetLogger("stuy-spec-uploader")

	// Log format string. Everything except the message has a custom color
	// which is dependent on the log level. Many fields have a custom output
	// formatting too, eg. the time returns the hour down to the milli second.
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)

	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter) // Set backends to be used

	return log
}
