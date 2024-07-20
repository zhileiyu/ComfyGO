package logger

import "log"

type logger interface {
	Debug(v ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

type defaultLogger struct{}

func SetLogger(l logger) {

}

func Debug(v ...interface{}) {
	log.Println(v)
}

func Info(v ...interface{}) {
	log.Println(v)
}

func Error(v ...interface{}) {
	log.Println(v)
}

func Fatal(v ...interface{}) {
	log.Fatal(v...)
}
