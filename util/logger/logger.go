// Package logger is the basic library of writing log.
// There are four level in logger, Error(0), Warn(1), Info(2) and Trace(3).
// If show level is set to 'n', only level smaller or equal to 'n' will
// print to standard output.
package logger

import (
	"fmt"
	"github.com/KenerChang/api-server/middleware"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	// Trace is for debug only
	Trace *Logger

	// Info is for normal log
	Info *Logger

	// Warn if for protential error
	Warn *Logger

	// Error if for critical error
	Error *Logger

	logLevel = map[string]int{
		"ERROR": 0,
		"WARN":  1,
		"INFO":  2,
		"TRACE": 3,
	}
	levelCount = 4
	logPrefix  = ""
	showLevel  = 1
)

const (
	levelError = iota
	levelWarn
	levelInfo
	levelTrace
)

type Logger struct {
	*log.Logger
}

func requestidFromContext(r *http.Request) string {
	ctx := r.Context()

	reqIDRaw := ctx.Value(middleware.ContextKeyRequestID) // reqIDRaw at this point is of type 'interface{}'

	reqID, ok := reqIDRaw.(string)
	if !ok {
		// handler error
		return ""
	}
	return reqID
}

func (l *Logger) Println(r *http.Request, v ...interface{}) {
	msg := fmt.Sprint(v...)
	if r == nil {
		l.Output(2, msg)
		return
	}
	rid := requestidFromContext(r)

	msg = fmt.Sprintf("[%s] %s", rid, msg)
	l.Output(2, msg)
}

func (l *Logger) Printf(r *http.Request, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	if r == nil {
		l.Output(2, msg)
		return
	}

	rid := requestidFromContext(r)

	msg = fmt.Sprintf("[%s] %s", rid, msg)
	l.Output(2, msg)
}

func init() {
	Init("", os.Stdout, os.Stdout, os.Stdout, ioutil.Discard)
}

// Init will init logger package with specific prefix and outputs.
// First parameter is prefix, and after second will be output of different level in order of:
// 	ERROR, WARN, INFO, TRACE.
// If parameter less then 5, level without output will use ioutil.Discard
func Init(
	prefix string,
	handler ...io.Writer) {
	logPrefix = prefix

	for len(handler) < levelCount {
		handler = append(handler, ioutil.Discard)
	}

	eLoger := log.New(handler[logLevel["ERROR"]],
		"[ERROR] ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Error = &Logger{
		eLoger,
	}

	wLoger := log.New(handler[logLevel["WARN"]],
		"[WARN] ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Warn = &Logger{
		wLoger,
	}

	iLoger := log.New(handler[logLevel["INFO"]],
		"[INFO] ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Info = &Logger{
		iLoger,
	}

	tLoger := log.New(handler[logLevel["TRACE"]],
		"[TRACE] ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Trace = &Logger{
		tLoger,
	}
}

// SetPrefix will set log start with [prefix]
func SetPrefix(prefix string) {
	logPrefix = prefix
	output := createOutputIO(showLevel)
	Init(logPrefix, output...)
}

// SetLevel will set minimum level output to stdout.
// Level can be one of "ERROR", "WARN", "INFO", "TRACE".
// If input is not one of above, level will set to INFO
func SetLevel(level string) {
	var ok bool
	showLevel, ok = logLevel[level]
	if !ok {
		showLevel = 1
	}
	output := createOutputIO(showLevel)
	Init(logPrefix, output...)
}

func createOutputIO(level int) []io.Writer {
	output := make([]io.Writer, levelCount)
	for idx := range output {
		if idx <= level {
			output[idx] = os.Stdout
		} else {
			output[idx] = ioutil.Discard
		}
	}
	return output
}
