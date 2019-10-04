package utils

import (
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

// HandleError saves error message to a log file if it is not nil
// returns true if there is an error
func HandleError(text string, err error) bool {
	if err == nil {
		return false
	}

	_, fileName, lineNum, _ := runtime.Caller(1)
	date := time.Now()

	dateInfo := "[" + date.Format("02-01-2006 15:04:05") + "] "
	msgError := fileName + ":" + strconv.Itoa(lineNum) + ": " + text + ". " + err.Error() + "\n"

	createLogFileIfDoesntExist("logs.log")
	logFile, err := os.OpenFile("logs.log", os.O_APPEND|os.O_WRONLY, 0600)

	printIfExist(err)

	log.Print(msgError)

	_, writeErr := logFile.WriteString(dateInfo + msgError)
	printIfExist(writeErr)

	printIfExist(logFile.Close())

	return true
}

func createLogFileIfDoesntExist(filePath string) {
	if !FileExists(filePath) {
		err := ioutil.WriteFile(filePath, []byte(""), 0600)
		HandleError("Can't create empty log file", err)
	}
}

func printIfExist(err error) {
	if err != nil {
		log.Print(err)
	}
}