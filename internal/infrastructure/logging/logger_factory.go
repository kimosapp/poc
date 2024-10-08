package logging

import (
	"errors"
	"github.com/kimosapp/poc/internal/core/ports/logging"
	"os"
)

func NewLogger() (logging.Logger, error) {
	loggerImplementation := os.Getenv("LOGGER_IMPLEMENTATION")
	if loggerImplementation == "" {
		if os.Getenv("ENV") == "dev" {
			return &StandardLogger{}, nil
		}
	}
	if loggerImplementation == "standard" {
		return &StandardLogger{}, nil
	}
	return nil, errors.New("invalid logger implementation")
}
