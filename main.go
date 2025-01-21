package logging

import (
	"fmt"
	"log"
	"os"
	"path"
)

func Master(show bool, enableCustomFileName bool, enableCustomLogLineNumber bool) {
	if show {
		customFileName(enableCustomFileName)
		customLogLineNumber(enableCustomLogLineNumber)
	}
}

func customFileName(enabled bool) {
	if enabled {
		LOGFILE := path.Join(os.TempDir(), "learning-go.log")
		fmt.Println("Log file:", LOGFILE)
		f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer f.Close()
		iLog := log.New(f, "iLog", log.LstdFlags)
		iLog.Println("Hello There!")
		iLog.Println("Mastering Go 4th Edition")
	}
}

func customLogLineNumber(enabled bool) {
	if enabled {
		LOGFILE := path.Join(os.TempDir(), "learning-go-line-number.log")
		fmt.Println("Log file:", LOGFILE)
		f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer f.Close()
		LstdFlags := log.Ldate | log.Lshortfile
		iLog := log.New(f, "LNum", LstdFlags)
		iLog.Println("Mastering Go 4th Edition")
		iLog.SetFlags(log.Lshortfile | log.LstdFlags)
		iLog.Println("Another Log Entry")
	}
}
