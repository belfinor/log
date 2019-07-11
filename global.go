package log

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2019-07-11

var defLog *Log

func SetLevel(level string) {
	defLog.SetLevel(level)
}

func GetLevel() string {
	return defLog.GetLevel()
}

func Trace(str ...interface{}) {
	defLog.Trace(str...)
}

func Debug(str ...interface{}) {
	defLog.Debug(str...)
}

func Warn(str ...interface{}) {
	defLog.Warn(str...)
}

func Info(str ...interface{}) {
	defLog.Info(str...)
}

func Error(str ...interface{}) {
	defLog.Error(str...)
}

func Finish(str ...interface{}) {
	defLog.Finish(str...)
}

func Fatal(str ...interface{}) {
	defLog.Fatal(str...)
}

func Logger(level string, strs ...interface{}) {
	defLog.Logger(level, strs)
}
