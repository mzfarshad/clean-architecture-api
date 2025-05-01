package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	apperr "github.com/mzfarshad/music_store_api/pkg/appErr"
)

type Option struct {
	RequestID        string
	ShowLine         bool
	ShowFunctionName bool
	ShowTimeStamp    bool
	ShowFileName     bool
}

type opt func(*Option)

func (l *logger) Conf(setup opt) {
	setup(l.Config)
}

func applyOption(requestId, level string, option Option) (*logPrint, error) {
	pc, file, line, ok := findCaller(2)
	if !ok {
		customErr := apperr.NewAppErr(
			apperr.StatusInternalServerError,
			"failed runtime.Caller(2) in logger pkg",
			apperr.TypeInternal,
			"",
		)
		return nil, customErr
	}
	logEty := &logPrint{}
	if option.ShowTimeStamp {
		logEty.TimeStamp = time.Now().Format(time.RFC3339)
	}
	logEty.Level = level
	if requestId != "unknown" {
		logEty.RequestID = requestId
	}
	if option.ShowFileName {
		logEty.FileName = filepath.Base(file)
	}
	if option.ShowFunctionName {
		fn := runtime.FuncForPC(pc).Name()
		logEty.FuncName = filepath.Base(fn)
	}
	if option.ShowLine {
		logEty.Line = fmt.Sprintf("Line: %d ", line)
	}
	return logEty, nil
}

// findCaller tries to identify the first caller outside of the logger package,
// starting from the provided skip value. It skips internal logger calls to
// accurately capture the original function that invoked the logger.
func findCaller(skip int) (uintptr, string, int, bool) {
	for i := skip; i < 20; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if !strings.Contains(file, "/logger/") && !strings.Contains(file, "logger") {
			return pc, file, line, ok
		}
	}
	return 0, "", 0, false
}
