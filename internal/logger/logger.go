package logger

import "log"

type logger interface {
	Debug(v ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

var userLogger logger

func SetLogger(l logger) {
	userLogger = l
}

func Debug(v ...interface{}) {
	if userLogger != nil {
		userLogger.Debug(v...)
		return
	}
	log.Printf("comfy client debug log: %v", v)
}

func Info(v ...interface{}) {
	if userLogger != nil {
		userLogger.Debug(v...)
		return
	}
	log.Printf("comfy client info log: %v", v)
}

func Error(v ...interface{}) {
	if userLogger != nil {
		userLogger.Debug(v...)
		return
	}
	log.Printf("comfy client error log: %v", v)
}

func Fatal(v ...interface{}) {
	if userLogger != nil {
		userLogger.Debug(v...)
		return
	}
	log.Fatal("comfy client fatal log: %v", v)
}
