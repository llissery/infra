package log

import (
	"github.com/rs/zerolog/log"
)

func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
