package loggdb

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/muesli/termenv"
	"github.com/snwfdhmp/errlog"
)

// Logger Main Logger Definition
type Logger struct {
	*log.Logger
	Options    *log.Options
	LogWriter  io.Writer
	LogOptions *CustomOpt
	LogDir          string
	Gdb        errlog.Logger
}

// CustomOpt Optional Options to pass to log
type CustomOpt struct {
	Prefix          string
	TimeFunction    func() time.Time
	TimeFormat      string
	ReportTimestamp bool
	ReportCaller    bool
	LogFileName     string
}

// Constants
const (
	Info  = log.InfoLevel
	Debug = log.DebugLevel
	Warn  = log.WarnLevel
	Error = log.ErrorLevel
	Fatal = log.FatalLevel
)

var (
	DebugPrint = func(format string, data ...interface{}) {
		fmt.Printf(format+"\n", data...)
	}
)

// newDebugger Creates a new debugger to backtrace errors.
func (Logger *Logger) newDebugger() {
	Config := &errlog.Config{
		PrintFunc:          DebugPrint,
		LinesBefore:        3,
		LinesAfter:         3,
		PrintError:         true,
		PrintSource:        true,
		PrintStack:         false,
		ExitOnDebugSuccess: false,
	}
	Logger.Gdb = errlog.NewLogger(Config)
}

// setOptions Set Default Options
func (Log *Logger) setOptions() {
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

// createLog Creates a Log file to save all output to
func (Log *Logger) createLog() error {
	var err error
	if Log.LogDir == "" {
		Log.LogDir, err = os.Getwd()
		if err != nil {
			return err
		}
	}
	if _, err = os.Stat(Log.LogDir); os.IsNotExist(err) {
		if err = os.Mkdir(Log.LogDir, 0744); err != nil {
			return err
		}
	}
	file := fmt.Sprintf("%s/%s", Log.LogDir, Log.LogOptions.LogFileName)
	LogFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	Log.LogWriter = io.MultiWriter(LogFile, os.Stdout)
	return nil
}

// NewLogger Creates a new Logger.
func (Log *Logger) NewLogger() error {
	Log.setOptions()
	Log.newDebugger()
	err := Log.createLog()
	if err != nil {
		return err
	}
	Logger := log.NewWithOptions(Log.LogWriter, *Log.Options)
	Log.Logger = Logger
	// Enforce ANSI256 to fixes colors.
	Log.Logger.SetColorProfile(termenv.ANSI256)
	return nil
}
