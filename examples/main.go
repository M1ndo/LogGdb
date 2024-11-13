package main

import (
	"errors"

	loggdb "github.com/m1ndo/LogGdb"
)

func main() {
	Log := &loggdb.Logger{}
	// CustomOpt := &loggdb.CustomOpt{ // Set a custom Options
	// 	Prefix: "Ran 911",
	// 	TimeFormat: time.Layout,
	// 	ReportTimestamp: true,
	// }
	// Logger.LogOptions = CustomOpt
	err := Log.NewLogger()
	if err != nil {
		panic(err)
	}
	Log.Info("This is a info")
	Log.Warn("This is a warning")
	Log.Debug("This is a Debug ( WILL NOT OUTPUT )")
	Log.Error("This is a Error")
	Log.SetLevel(loggdb.Debug)
	Log.Debug("This is a debug (WILL OUTPUT)")
	Log.Gdb.Debug(errors.New("This is a serious error"))
	Log.Info("Will this print?")
	LogLevel := Log.GetLevel()
	Log.Infof("Current Love Level %s",  LogLevel)
	Log.Fatal("Simply Die (return 1 err code)")
}
