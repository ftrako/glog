package glog

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// log.Println()接口二次封装，接口更丰富

type Logger struct {
	Color     bool  `json:"color"`      // 是否开启字体颜色
	Prefix    bool  `json:"prefix"`     // 是否开启前缀，比如[INFO]
	Func      bool  `json:"func"`       // 是否显示函数名
	Flag      int   `json:"flag"`       // 属性，参考log.LstdFlags
	MinLevel  Level `json:"min_level"`  // 最低日志级别，参考Level类型
	CallDepth int   `json:"call_depth"` // 调用函数深度，特意提供可支持外部修改
}

func NewLogger() *Logger {
	log.SetFlags(0) // 不带任何属性，不用原系统属性
	b := new(Logger)
	b.Color = false
	b.Prefix = true
	b.Func = true
	b.MinLevel = LevelDebug
	b.Flag = log.Lmicroseconds | log.Lshortfile
	b.CallDepth = 3
	return b
}

func (p *Logger) clone() *Logger {
	return &Logger{
		Color:     p.Color,
		Prefix:    p.Prefix,
		Func:      p.Func,
		MinLevel:  p.MinLevel,
		Flag:      p.Flag,
		CallDepth: p.CallDepth,
	}
}

// ln - true表示自动换行
func (p *Logger) D(f interface{}, v ...interface{}) {
	p.write(LevelDebug, f, v...)
}

func (p *Logger) I(f interface{}, v ...interface{}) {
	p.write(LevelInfo, f, v...)
}

func (p *Logger) W(f interface{}, v ...interface{}) {
	p.write(LevelWarn, f, v...)
}

func (p *Logger) E(f interface{}, v ...interface{}) {
	p.write(LevelError, f, v...)
}

func (p *Logger) write(level Level, f interface{}, v ...interface{}) {
	if level < p.MinLevel { // 级别限制
		return
	}
	now := time.Now()

	caller := ""
	if p.Flag&(log.Lshortfile|log.Llongfile) != 0 {
		var ok bool
		var funcName = ""
		fnc, file, line, ok := runtime.Caller(p.CallDepth)
		if !ok {
			file = "???"
			line = 0
		} else {
			if p.Flag&log.Lshortfile != 0 {
				short := file
				for i := len(file) - 1; i > 0; i-- {
					if file[i] == '/' {
						short = file[i+1:]
						break
					}
				}
				file = short
			}
			if p.Func {
				// 去掉.之前的文件名
				short := runtime.FuncForPC(fnc).Name()
				for i := len(short) - 1; i > 0; i-- {
					if short[i] == '.' {
						short = short[i+1:]
						break
					}
				}
				funcName = " " + short
			}
		}
		caller = "[" + file + ":" + strconv.Itoa(line) + funcName + "] "
	}

	t := ""
	if p.Flag&log.Lmicroseconds != 0 {
		t = now.Format("2006-01-02 15:04:05.000") + " "
	} else if p.Flag&log.Ltime != 0 {
		t = now.Format("2006-01-02 15:04:05") + " "
	} else if p.Flag&log.Ldate != 0 {
		t = now.Format("2006-01-02") + " "
	}
	prefix := ""
	if p.Prefix {
		prefix = "[" + StrLevelMap[level] + "] "
	}

	str := t + prefix + caller + p.format(f, v...)
	if p.Color {
		str = colors[level](str)
	}

	if level < LevelError {
		fmt.Println(str)
	} else {
		_, _ = os.Stderr.WriteString(str)
	}
}

// f - format
// v - 系列参数
func (p *Logger) format(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			// format string
		} else {
			// do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}
