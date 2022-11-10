package logger

import (
	"log"
	"os"
)

type Logger struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func New() Logger {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return Logger{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}
}
