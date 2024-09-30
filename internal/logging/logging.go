package logging

import (
	"io"
	"os"
)

func InitLogOutput(logfile string, multiwrite bool) io.Writer {
	file, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	if multiwrite {
		return io.MultiWriter(file, os.Stdout)
	}
	return file
}
