package lib

import (
	"os"

	"github.com/go-logr/logr"
)

func AssertErr(err error, message string, logger logr.Logger) {
	if err != nil {
		logger.Error(err, message)
		os.Exit(1)
	}
}
