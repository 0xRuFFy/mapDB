package server

import (
	"log"
)

type MapDBLogger struct {
	logger *log.Logger
}

func newMapDBLogger() *MapDBLogger {
	return &MapDBLogger{
		logger: log.New(log.Writer(), "[Server] ", log.LstdFlags),
	}
}

var logger *MapDBLogger = newMapDBLogger()

func (l *MapDBLogger) Info(message string) {
	l.logger.Println("[INFO]", message)
}

func (l *MapDBLogger) Error(message string) {
	l.logger.Println("[ERROR]", message)
}

func (l *MapDBLogger) Fatal(message string) {
	l.logger.Fatalln("[FATAL]", message)
}

func (l *MapDBLogger) Panic(message string) {
	l.logger.Panicln("[PANIC]", message)
}

func (l *MapDBLogger) Debug(message string) {
	l.logger.Println("[DEBUG]", message)
}

func (l *MapDBLogger) Trace(message string) {
	l.logger.Println("[TRACE]", message)
}

func (l *MapDBLogger) Warn(message string) {
	l.logger.Println("[WARN]", message)
}
