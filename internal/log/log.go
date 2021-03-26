// package log aims to provide common logging interface for
// the CLI.
package log

import (
	"fmt"
)

func Info(message string) {
	fmt.Printf("[info] %v\n", message)
}

func Infof(message string, args ...interface{}) {
	fmt.Printf("[info] "+message+"\n", args...)
}

func Error(message string) {
	fmt.Printf("[error] %v\n", message)
}

func Errorf(message string, args ...interface{}) {
	fmt.Printf("[error] "+message, args...)
}

func Debug(message string) {
	fmt.Printf("[debug] %v\n", message)
}

func Warn(message string) {
	fmt.Printf("[warning] %v\n", message)
}

func Warnf(message string, args ...interface{}) {
	fmt.Printf("[warning] "+message, args...)
}
