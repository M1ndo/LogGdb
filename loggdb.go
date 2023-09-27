package loggdb

import (
	"io"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

type Logger struct {
	Log        *log.Logger
	Options    *log.Options
	LogWriter  io.Writer
	LogDir     string
	LogOptions *CustomOpt
}

type CustomOpt struct {
	Prefix          string
	TimeFunction    func() time.Time
	TimeFormat      string
	ReportTimestamp bool
	ReportCaller    bool
}

const (
	Info  = log.InfoLevel
	Debug = log.DebugLevel
	Warn  = log.WarnLevel
	Error = log.ErrorLevel
	Fatal = log.FatalLevel
)

func (Log *Logger) SetOptions() {
	if Log.LogOptions != nil {
		Log.Options = &log.Options{
			Prefix:          Log.LogOptions.Prefix,
			TimeFunction:    Log.LogOptions.TimeFunction,
			TimeFormat:      Log.LogOptions.TimeFormat,
			ReportTimestamp: Log.LogOptions.ReportTimestamp,
			ReportCaller:    Log.LogOptions.ReportCaller,
		}
		return
	}
	Log.Options = &log.Options{
		TimeFunction:    time.Now,
		TimeFormat:      time.DateTime,
		ReportTimestamp: true,
		Level:           Info,
		Prefix:          "Logging 👾 ",
	}
}

func (Log *Logger) createLog() error {
	if Log.LogDir == "" {
		var err error
		Log.LogDir, err = os.Getwd()
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(Log.LogDir); os.IsNotExist(err) {
		if err = os.Mkdir(Log.LogDir, 0744); err != nil {
			return err
		}
	}
	file := Log.LogDir + "/logger.log"
	LogFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	Log.LogWriter = io.MultiWriter(LogFile, os.Stdout)
	return nil
}

func (Log *Logger) NewLogger() error {
	Log.SetOptions()
	err := Log.createLog()
	if err != nil {
		return err
	}
	Logger := log.NewWithOptions(Log.LogWriter, *Log.Options)
	Log.Log = Logger
	return nil
}
