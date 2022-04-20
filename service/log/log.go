package log

import (
	"fmt"
	"github.com/wtifs/ddmc/constants"
	"time"
)

func Debug(format string, msg ...any) {
	Log("DEBUG", format, msg...)
}

func Err(format string, msg ...any) {
	Log("ERROR", format, msg...)
}

func Info(format string, msg ...any) {
	Log("INFO", format, msg...)
}

func Log(level, format string, msg ...any) {
	fmt.Printf("[%s] %s ", level, time.Now().Format(constants.DateTimeFmt))
	fmt.Printf(format+"\n", msg...)
}
