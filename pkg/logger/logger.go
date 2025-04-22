package logger

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	apperr "github.com/mzfarshad/music_store_api/pkg/appErr"
)

type Logger struct {
	LogFile  *os.File
	LogLevel string
	Mutex    *sync.Mutex
	Writer   *bufio.Writer
}

type contextKey string

const (
	loggerKey contextKey = "logger"
)

func WithLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func GetLogger(ctx context.Context) *Logger {
	if logger, ok := ctx.Value(loggerKey).(*Logger); ok {
		return logger
	}
	return nil
}

func NewLogger(logLevel string) (*Logger, error) {
	// check logs/ exists or no
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		if err := os.Mkdir("logs", 0755); err != nil {
			customErr := apperr.NewAppErr(
				apperr.StatusInternalServerError,
				"faield create logs directory",
				apperr.TypeInternal,
				err.Error(),
			)
			return nil, customErr
		}
	}
	filePath := "logs/app.log"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		customErr := apperr.NewAppErr(
			apperr.StatusInternalServerError,
			"failed open app.log file",
			apperr.TypeInternal,
			err.Error(),
		)
		return nil, customErr
	}
	writer := bufio.NewWriter(file)
	return &Logger{
		LogFile:  file,
		LogLevel: logLevel,
		Mutex:    &sync.Mutex{},
		Writer:   writer,
	}, nil
}

func (l *Logger) Log(ctx context.Context, level, message string, err error) {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	requestID := "-"
	if reqID := ctx.Value("requestID"); reqID != nil {
		requestID = reqID.(string)
	}

	timeStamp := time.Now().Format("2006:01:02 15:04:05")

	var logMessage string

	if err != nil {
		if customErr, ok := err.(*apperr.CustomErr); ok {
			logMessage = fmt.Sprintf("[%s] [%s] [requestID: %s]  %s | %s | %s\n",
				timeStamp, level, requestID, message, customErr.Message, customErr.Details)
		} else {
			logMessage = fmt.Sprintf("[%s] [%s] [requestID: %s]  %s | %s\n",
				timeStamp, level, requestID, message, err.Error())
		}
	} else {
		logMessage = fmt.Sprintf("[%s] [%s] [requestID: %s]  %s\n",
			timeStamp, level, requestID, message)
	}

	if _, err := l.Writer.WriteString(logMessage); err != nil {
		customErr := apperr.NewAppErr(
			apperr.StatusInternalServerError,
			"failed write to bufio Log function in logger package",
			apperr.TypeInternal,
			err.Error(),
		)
		log.Println(customErr)
		return
	}
	if err := l.Writer.Flush(); err != nil {
		customErr := apperr.NewAppErr(
			apperr.StatusInternalServerError,
			"failed flush log to logs/app.log file",
			apperr.TypeInternal,
			err.Error(),
		)
		log.Println(customErr)
		return
	}
}

func (l *Logger) Debug(ctx context.Context, message string, err error) {
	l.Log(ctx, "DEBUG", message, err)
}

func (l *Logger) Info(ctx context.Context, message string, err error) {
	l.Log(ctx, "INFO", message, err)
}

func (l *Logger) Warn(ctx context.Context, message string, err error) {
	l.Log(ctx, "WARN", message, err)
}

func (l *Logger) Fatal(ctx context.Context, message string, err error) {
	l.Log(ctx, "FATAL", message, err)
	os.Exit(1)
}

func (l *Logger) Panic(ctx context.Context, message string, err error) {
	l.Log(ctx, "PANIC", message, err)
	panic(err)
}
