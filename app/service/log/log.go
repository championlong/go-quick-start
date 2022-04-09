package log

import (
	"fmt"
)

func Debug(format string, v ...interface{}) {
	fmt.Printf("[DEBUG] "+format+"\n", v...)
}

func Info(format string, v ...interface{}) {
	fmt.Printf("[INFO] "+format+"\n", v...)
}

func Warning(format string, v ...interface{}) {
	fmt.Printf("[WARNING] "+format+"\n", v...)
}

func Err(format string, v ...interface{}) {
	fmt.Printf("[ERR] "+format+"\n", v...)
}

// 重大错误
func Fatal(format string, v ...interface{}) {
	fmt.Printf("[FATAL] "+format+"\n", v...)
	panic(fmt.Sprintf(format, v...))
}

func Println(v interface{}) {
	println("%+v", v)
}
