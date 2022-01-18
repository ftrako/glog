package glog

var defaultLogger *Logger

func init() {
	defaultLogger = NewLogger()
}

// 参考 log.LstdFlags
func SetFlag(flag int) {
	defaultLogger.Flag = flag
}

// 是否开启字体颜色
func EnableColor(enable bool) {
	defaultLogger.Color = enable
}

// 是否开启前缀，比如[INFO]
func EnablePrefix(enable bool) {
	defaultLogger.Prefix = enable
}

// 是否显示函数名
func EnableFuncName(enable bool) {
	defaultLogger.Func = enable
}

// 设置最低日志级别，低于该级别的将不会打印
// level值参考 LevelDebug
func SetMinLevel(minLevel Level) {
	defaultLogger.MinLevel = minLevel
}

// 设置函数调用层级，从级别0开始
func FuncDepth(depth int) *Logger {
	if depth < 0 {
		depth = 0
	}
	l := defaultLogger.clone()
	l.CallDepth = depth + 2
	return l
}

func D(f interface{}, v ...interface{}) {
	FuncDepth(1).D(f, v...)
}

func I(f interface{}, v ...interface{}) {
	FuncDepth(1).I(f, v...)
}

func W(f interface{}, v ...interface{}) {
	FuncDepth(1).W(f, v...)
}

func E(f interface{}, v ...interface{}) {
	FuncDepth(1).E(f, v...)
}
