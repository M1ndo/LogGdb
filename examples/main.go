package main

import (
	loggdb "github.com/m1ndo/LogGdb"
)

func main() {
	Logger := &loggdb.Logger{}
	// CustomOpt := &loggdb.CustomOpt{ // Set a custom Options
	// 	Prefix: "Ran 911",
	// 	TimeFormat: time.Layout,
	// 	ReportTimestamp: true,
	// }
	// Logger.LogOptions = CustomOpt
	err := Logger.NewLogger()
	if err != nil {
		panic(err)
	}
	Logger.Log.Info("This is a info")
	Logger.Log.Warn("This is a warning")
	Logger.Log.Debug("This is a Debug ( WILL NOT OUTPUT )")
	Logger.Log.Error("This is a Error")
	Logger.Log.SetLevel(loggdb.Debug)
	Logger.Log.Debug("This is a debug (WILL OUTPUT)")
	Logger.Log.Info("Will this print?")
	LogLevel := Logger.Log.GetLevel()
	Logger.Log.Infof("Current Love Level %s",  LogLevel)
	Logger.Log.Fatal("Simply Die (return 1 err code)")
}
