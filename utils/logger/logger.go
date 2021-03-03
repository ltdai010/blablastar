package logger

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"log"
	"sync"
)

var (
	_logger *logs.BeeLogger
	once    sync.Once
)

func init() {
	once.Do(func() {
		onceInitLogger()
	})
}

func onceInitLogger() {
	fmt.Println("init logger")
	_logger = logs.NewLogger()
	err := _logger.SetLogger(logs.AdapterFile, `{"filename":"logs/BlaBla.log","maxdays": 365,"perm":"0744"}`)
	if err != nil {
		log.Fatal(err)
	}
}

// Debug write log in debug level with format
func Debug(format string, v ...interface{}) {
	_logger.Debug(format, v...)
	log.Printf("[Debug] "+format+"\n", v...)
}

// Warn write log in warning level with format
func Warn(format string, v ...interface{}) {
	_logger.Warn(format, v...)
	log.Printf("[Warn] "+format+"\n", v...)
}

// Error write log in error level with format
func Error(format string, v ...interface{}) {
	_logger.Error(format, v...)
	log.Printf("[Error] "+format+"\n", v...)
}

// Info write log in info level with format
func Info(format string, v ...interface{}) {
	_logger.Info(fmt.Sprintf(format, v...))
	log.Printf("[Info] "+format+"\n", v...)
}

