package logger

import (
	"bufio"
	"io"
	"log"
	"os"

	apperr "github.com/mzfarshad/music_store_api/pkg/appErr"
)

type logPrint struct {
	TimeStamp    string `json:"time_stamp,omitempty"`
	Level        string `json:"level"`
	RequsetID    string `json:"request_id,omitempty"`
	FileName     string `json:"file_name,omitempty"`
	FuncName     string `json:"function_name,omitempty"`
	Line         string `json:"line,omitempty"`
	Message      string `json:"message,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	ErrorDetails string `json:"error_details,omitempty"`
}

func writeToFile(data string, writer *bufio.Writer, alsoStdout bool) {
	var outputWriter *bufio.Writer
	if alsoStdout {
		multi := io.MultiWriter(writer, os.Stdout)
		outputWriter = bufio.NewWriter(multi)
	} else {
		outputWriter = writer
	}

	if _, err := outputWriter.WriteString(data); err != nil {
		customErr := apperr.NewAppErr(
			apperr.StatusInternalServerError,
			"failed write to bufio Log function in logger package",
			apperr.TypeInternal,
			err.Error(),
		)
		log.Fatal(customErr)
	}
	if err := outputWriter.Flush(); err != nil {
		customErr := apperr.NewAppErr(
			apperr.StatusInternalServerError,
			"failed flush log to logs/app.log file",
			apperr.TypeInternal,
			err.Error(),
		)
		log.Fatal(customErr)
	}
}
