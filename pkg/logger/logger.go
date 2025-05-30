package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

func GetWorkDir() string {
	dir, err := os.Getwd()
	if err != nil {
		ErrorHandler(err, "logger/GetWorkDir/")
	}
	return dir
}

func ErrorHandler(err error, where string) {
	varErr := fmt.Sprint(err) + "\nМесто: " + where
	WriteToLogger("ERROR: " + varErr)
	log.Fatalln(varErr)
}

func InfoHandler(err error, where string) {
	varErr := fmt.Sprint(err) + "\nМесто: " + where
	WriteToLogger("WARN: " + varErr)
	log.Println(varErr)
}

func WriteToLogger(msg string) {
	os.MkdirAll("logs", os.ModePerm)
	file, err := os.Create("log_" + time.Now().Local().Format("2006-01-02"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.WriteString(msg)
	if err != nil {
		log.Fatal(err)
	}
}
