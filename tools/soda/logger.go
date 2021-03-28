package soda

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/gobuffalo/pop/v5/logging"
)

type Logger struct{}

func (l *Logger) Log(lvl logging.Level, s string, args ...interface{}) {
	if lvl == logging.Debug {
		return
	}

	if lvl == logging.SQL {
		return
	}

	s = fmt.Sprintf(s, args...)
	s = fmt.Sprintf("%s - %s", lvl, s)
	s = color.YellowString(s)

	log := log.New(os.Stdout, "[POP] ", log.LstdFlags)

	log.Println(s)
}
