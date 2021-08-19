package logger

var defaultLogger *Logger

// 参考log.LstdFlags
func SetFlag(flag int) {
    defaultLogger.Flag = flag
}

func EnableColor(enable bool) {
    defaultLogger.Color = enable
}

func EnablePrefix(enable bool) {
    defaultLogger.Prefix = enable
}

func EnableFuncName(enable bool) {
    defaultLogger.Func = enable
}

// 设置日志级别
// level值参考 LevelDebug
func SetLevel(level int) {
    defaultLogger.Level = level
}

func Trace(f interface{}, v ...interface{}) {
    TraceDepth(1,f,v...)
}

// depth层级数
func TraceDepth(depth int, f interface{}, v ...interface{}) {
    defaultLogger.TraceDepth(depth+3, f, v...)
}

func Debug(f interface{}, v ...interface{}) {
    DebugDepth(1,f,v...)
}

// depth层级数
func DebugDepth(depth int, f interface{}, v ...interface{}) {
    defaultLogger.DebugDepth(depth+3, f, v...)
}

func Info(f interface{}, v ...interface{}) {
    InfoDepth(1,f, v...)
}

func InfoDepth(depth int, f interface{}, v ...interface{}) {
    defaultLogger.InfoDepth(depth+3, f, v...)
}

func Warn(f interface{}, v ...interface{}) {
    WarnDepth(1,f, v...)
}

func WarnDepth(depth int, f interface{}, v ...interface{}) {
    defaultLogger.WarnDepth(depth+3, f, v...)
}

func Error(f interface{}, v ...interface{}) {
    ErrorDepth(1,f, v...)
}

func ErrorDepth(depth int, f interface{}, v ...interface{}) {
    defaultLogger.ErrorDepth(depth+3, f, v...)
}

func Panic(f interface{}, v ...interface{}) {
    PanicDepth(1,f, v...)
}

func PanicDepth(depth int, f interface{}, v ...interface{}) {
    defaultLogger.PanicDepth(depth+3, f, v...)
}

// Deprecated
func GetDefaultLogger() *Logger {
    return defaultLogger
}

func init() {
    defaultLogger = NewLogger()
}
