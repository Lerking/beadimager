package main

import (
	"fmt"
	"log"
	"os"
)

var LoggingFile *os.File

func InitLogFile() {
	_, err := os.Stat(fmt.Sprintf(UserPath + Sep + LogFile))
	if err == nil {
		LoggingFile, err := os.OpenFile(fmt.Sprintf(UserPath+Sep+LogFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(LoggingFile)
	} else {
		if _, err := os.Stat(UserPath); err != nil {
			os.Mkdir(UserPath, 0755)
		}
		_, err = os.Create(fmt.Sprintf(UserPath + Sep + LogFile))
		if err != nil {
			log.Fatal(err)
		}
		LoggingFile, err := os.OpenFile(fmt.Sprintf(UserPath+Sep+LogFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(LoggingFile)
	}
	log.Print("Logging initialized")
}
