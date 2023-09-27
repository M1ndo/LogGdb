package loggdb

import (
	"io"
	"os"

	"github.com/charmbracelet/log"
)

type Logger struct {
	Log *log.Logger
	Options *log.Options
	LogWriter io.Writer
	LogDir string
}

func (Log *Logger) SetOptions() {
	Options := &log.Options{
		Level: log.InfoLevel,
		Prefix: "Logging 👾 ",
	}
	Log.Options = Options
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
	// newWriter := io.MultiWriter(os.Stdout)
	Log.LogWriter = io.MultiWriter(LogFile, os.Stdout)
	// Log.LogWriter = newWriter
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
