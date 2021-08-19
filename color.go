package logger

type brush func(string) string

func newBrush(color string) brush {
    pre := "\033["
    reset := "\033[0m"
    return func(text string) string {
        return pre + color + "m" + text + reset
    }
}

// 字体颜色：30:黑 31:红 32:绿 33:黄 34:蓝色 35:紫色 36:深绿 37:白色
// 背景：40:黑 41:深红 42:绿 43:黄色 44:蓝色 45:紫色 46:深绿 47:白色
var colors = []brush{
    newBrush("1;37"), // Debug              白色
    newBrush("1;37"), // Debug              白色
    newBrush("1;36"), // Info               天蓝色
    newBrush("1;33"), // Warn               黄色
    newBrush("1;31"), // Error              红色
    newBrush("1;41"), // Fatal              红色底
}
