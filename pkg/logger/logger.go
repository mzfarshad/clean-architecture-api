package logger

import (
	"bufio"
	"context"
	"encoding/json"
	"log"
	"os"
	"sync"

	apperr "github.com/mzfarshad/music_store_api/pkg/appErr"
)

type logger struct {
	LogFile  *os.File
	LogLevel string
	Mutex    *sync.Mutex
	Writer   *bufio.Writer
	Config   *Option
}

type contextKey string

const (
	loggerKey contextKey = "logger"
)

func WithLogger(ctx context.Context, logger *logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func GetLogger(ctx context.Context) *logger {
	if logger, ok := ctx.Value(loggerKey).(*logger); ok {
		return logger
	}
	return nil
}

func newLogger(logLevel string) (*logger, error) {
	// checking for existance of the log directory
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

	// default values in set up configer
	opt := &Option{
		RequestID:        "unknown",
		ShowLine:         false,
		ShowFunctionName: false,
		ShowTimeStamp:    true,
		ShowFileName:     false,
	}

	writer := bufio.NewWriter(file)
	return &logger{
		LogFile:  file,
		LogLevel: logLevel,
		Mutex:    &sync.Mutex{},
		Writer:   writer,
		Config:   opt,
	}, nil
}

var once sync.Once
var instance *logger

func Get(level string) *logger {
	once.Do(
		func() {
			l, err := newLogger(level)
			if err != nil {
				customeErr := apperr.NewAppErr(
					apperr.StatusInternalServerError,
					"faield make logger",
					apperr.TypeInternal,
					err.Error(),
				)
				log.Fatal(customeErr)
			}
			instance = l
		})
	return instance
}

func (l *logger) Log(ctx context.Context, level, message string, err error) {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	requestID := l.Config.RequestID
	if reqID := ctx.Value("requestID"); reqID != nil {
		requestID = reqID.(string)
	}
	logData, err := applyOption(requestID, level, *l.Config)
	if err != nil {
		customErr := apperr.NewAppErr(
			apperr.StatusForbidden,
			"failed set option into logger",
			apperr.TypeInternal,
			err.Error(),
		)
		log.Fatal(customErr)
	}

	if err != nil {
		if customErr, ok := err.(*apperr.CustomErr); ok {
			logData.Message = message
			logData.ErrorMessage = customErr.Message
			logData.ErrorDetails = customErr.Details

		} else {
			logData.Message = message
			logData.ErrorMessage = err.Error()
		}
	} else {
		logData.Message = message
	}

	jsonLogData, err := json.MarshalIndent(&logData, "", "  ")
	if err != nil {
		customErr := apperr.NewAppErr(
			apperr.StatusInternalServerError,
			"failed encode log data in logger package",
			apperr.TypeInternal,
			err.Error(),
		)
		log.Fatal(customErr)
	}
	strLogData := string(jsonLogData) + "\n"
	writeToFile(strLogData, l.Writer)
}

func (l *logger) Debug(ctx context.Context, message string, err error) {
	l.Log(ctx, "DEBUG", message, err)
}

func (l *logger) Info(ctx context.Context, message string, err error) {
	l.Log(ctx, "INFO", message, err)
}

func (l *logger) Warn(ctx context.Context, message string, err error) {
	l.Log(ctx, "WARN", message, err)
}

func (l *logger) Error(ctx context.Context, message string, err error) {
	l.Log(ctx, "ERROR", message, err)
}

func (l *logger) Fatal(ctx context.Context, message string, err error) {
	l.Log(ctx, "FATAL", message, err)
	os.Exit(1)
}

func (l *logger) Panic(ctx context.Context, message string, err error) {
	l.Log(ctx, "PANIC", message, err)
	panic(err)
}
