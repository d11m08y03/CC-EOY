package logger

import (
	"log"
	"os"
)

var f *os.File

func StartFileLogging() {
	f, err := os.OpenFile("logs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(f)
}

func StopFileLogging() {
	if f != nil {
		f.Close()
	}
}

func Info(message string) {
	log.Printf("[INFO]: %s", message)
}

func Warn(message string) {
	log.Printf("[WARN]: %s", message)
}

func Error(message string) {
	log.Printf("[ERROR]: %s", message)
}

func Fatal(message string) {
	log.Fatalf("[FATAL]: %s", message)
}
