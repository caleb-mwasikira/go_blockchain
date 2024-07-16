package utils

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

/*
This is a function that takes in a filename and returns
an io.Writer for a log file located in common log dir
*/
func CreateLogger(fpath string) io.Writer {
	fpath = filepath.Base(fpath)
	var defaultLogOutput io.Writer = os.Stdout

	// create/open log file
	logFpath := filepath.Join(LogDir, fpath)
	logFd, err := os.Create(logFpath)
	if err != nil {
		log.Println("error opening log file: ", err)
		return defaultLogOutput
	}

	logOutput := io.MultiWriter(defaultLogOutput, logFd)
	return logOutput
}
