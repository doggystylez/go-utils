package log

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

// Level represents the severity of the log message.
type Level int

// Log Levels
const (
	// none is a non-existent log level, used for configuration.
	none Level = iota
	Debug
	Info
	Error
)

// String representations of log levels
var levelStrings = []string{
	Debug: "[DBG]",
	Info:  "[INF]",
	Error: "[ERR]",
}

// ANSI color codes for log levels
var levelColors = []string{
	Debug: "[\033[33mDBG\033[0m]",
	Info:  "[\033[32mINF\033[0m]",
	Error: "[\033[31mERR\033[0m]",
}

const timeFormat = "2006/01/02 15:04:05"

// Logger is a simple logger that is safe for concurrent use.
type Logger struct {
	level   Level
	color   bool
	output  io.Writer
	mu      sync.RWMutex
	logChan chan Log
	sb      strings.Builder
	done    chan struct{}
	wg      sync.WaitGroup
}

// Log represents a log entry.
type Log struct {
	time  time.Time
	level Level
	label string
	msg   []any
}

// Opts defines the options for a logger.
type Opts struct {
	Level     string
	Color     bool
	BufferLen int
	Output    io.Writer
}

// ParseString parses a string into a log level.
func ParseString(level string) Level {
	switch strings.ToLower(level) {
	case "debug":
		return Debug
	case "info":
		return Info
	case "error":
		return Error
	default:
		return Info
	}
}

// NewLogger creates a new logger with the specified options. writer defaults to os.Stdout. level defaults to Info.
func NewLogger(opts *Opts) *Logger {
	logger := &Logger{
		level:   Info,
		output:  os.Stdout,
		logChan: make(chan Log, 100),
		done:    make(chan struct{}),
	}
	if opts != nil {
		if opts.Level != "" {
			logger.level = ParseString(opts.Level)
		}
		logger.color = opts.Color
		if opts.BufferLen != 0 {
			logger.logChan = make(chan Log, opts.BufferLen)
		}
		if opts.Output != nil {
			logger.output = opts.Output
		}
	}
	logger.wg.Add(1)
	go logger.processLogs()
	return logger
}

// processLogs handles log messages from the channel in its own goroutine.
func (l *Logger) processLogs() {
	defer l.wg.Done()
	for {
		select {
		case entry := <-l.logChan:
			l.write(entry)
		case <-l.done:
			for entry := range l.logChan {
				l.write(entry)
			}
			return
		}
	}
}

// write outputs the log entry to the writer.
func (l *Logger) write(entry Log) {
	l.mu.Lock()
	l.sb.WriteString(entry.time.Format(timeFormat))
	l.sb.WriteString(" ")
	if l.color {
		l.sb.WriteString(levelColors[entry.level])
	} else {
		l.sb.WriteString(levelStrings[entry.level])
	}
	l.sb.WriteString(" [")
	l.sb.WriteString(entry.label)
	l.sb.WriteString("] ")
	for _, m := range entry.msg {
		l.sb.WriteString(fmt.Sprint(m))
		l.sb.WriteString(" ")
	}
	fmt.Fprintln(l.output, strings.TrimSpace(l.sb.String()))
	l.sb.Reset()
	l.mu.Unlock()
}

// Debug logs a debug message.
func (l *Logger) Debug(label string, msg ...any) {
	l.log(Debug, label, msg...)
}

// Info logs an info message.
func (l *Logger) Info(label string, msg ...any) {
	l.log(Info, label, msg...)
}

// Error logs an error message.
func (l *Logger) Error(label string, msg ...any) {
	l.log(Error, label, msg...)
}

// log checks the log level and enqueues the log message if appropriate.
func (l *Logger) log(level Level, label string, msg ...any) {
	l.mu.RLock()
	levelOK := level >= l.level
	l.mu.RUnlock()
	if !levelOK {
		return
	}
	select {
	case l.logChan <- Log{time.Now().UTC(), level, label, msg}:
	default:
		fmt.Fprintf(os.Stderr, "Logger buffer is full, dropping log message: %v\n", msg)
	}
}

// Shutdown waits for the log queue to be processed and ceases logging.
func (l *Logger) Shutdown() {
	l.done <- struct{}{}
	close(l.done)
	close(l.logChan)
	l.wg.Wait()
}
