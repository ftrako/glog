package logger

type Level int

const (
    LevelTrace = iota
    LevelDebug
    LevelInfo
    LevelWarn
    LevelError
    LevelFatal
)

// 日志等级和描述映射关系
var LevelMap = map[string]int{
    "FATA": LevelFatal,
    "ERRO": LevelError,
    "WARN": LevelWarn,
    "INFO": LevelInfo,
    "DEBU": LevelDebug,
    "TRAC": LevelTrace,
}

var StrLevelMap = map[Level]string{
    LevelFatal: "FATA",
    LevelError: "ERRO",
    LevelWarn:  "WARN",
    LevelInfo:  "INFO",
    LevelDebug: "DEBU",
    LevelTrace: "TRAC",
}
