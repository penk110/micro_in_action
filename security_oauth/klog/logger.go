package klog

import (
	"log"
	"os"

	kitLog "github.com/go-kit/kit/log"
)

var (
	Logger    *log.Logger
	KitLogger kitLog.Logger
)

func init() {
	Logger = log.New(os.Stdout, "", log.LstdFlags)

	KitLogger = kitLog.NewLogfmtLogger(os.Stdout)
	KitLogger = kitLog.With(KitLogger, "ts", kitLog.DefaultTimestampUTC)
	KitLogger = kitLog.With(KitLogger, "caller", kitLog.DefaultCaller)

	Logger.Printf("init logger done!")
}
