package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
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
	l.Mutex.Lock()
	defer l.Mutex.Unlock()
	setup(l.Config)
}

func applyOption(requestId, level string, option Option) (*logPrint, error) {
	pc, file, line, ok := runtime.Caller(2)
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
		logEty.RequsetID = requestId
	}
	if option.ShowFileName {
		logEty.FileName = filepath.Base(file)
	}
	if option.ShowFunctionName {
		fn := runtime.FuncForPC(pc).Name()
		logEty.FuncName = filepath.Base(fn)
	}
	if option.ShowLine {
		logEty.Line = fmt.Sprintf("[Line: %d] ", line)
	}

	return logEty, nil
}
