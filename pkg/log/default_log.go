package log

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"strings"
	"sync"
	"time"
)

var (
	defLogger = New(NewZapOptions())
	mu        sync.Mutex
)

func Init(opts *ZapOptions) {
	mu.Lock()
	defer mu.Unlock()
	defLogger = New(opts)
}

func Flush() {
	_ = defLogger.logger.Sync()
}

func Debug(msg string, fields ...Field) {
	defLogger.logger.Debug(msg, fields...)
}

func Debugf(format string, v ...interface{}) {
	defLogger.logger.Sugar().Debugf(format, v...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	defLogger.logger.Sugar().Debugw(msg, keysAndValues...)
}

func Info(msg string, fields ...Field) {
	defLogger.logger.Info(msg, fields...)
}

func Infof(format string, v ...interface{}) {
	defLogger.logger.Sugar().Infof(format, v...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	defLogger.logger.Sugar().Infow(msg, keysAndValues...)
}

func Warn(msg string, fields ...Field) {
	defLogger.logger.Warn(msg, fields...)
}

func Warnf(format string, v ...interface{}) {
	defLogger.logger.Sugar().Warnf(format, v...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	defLogger.logger.Sugar().Warnw(msg, keysAndValues...)
}

func Error(msg string, fields ...Field) {
	defLogger.logger.Error(msg, fields...)
}

func Errorf(format string, v ...interface{}) {
	defLogger.logger.Sugar().Errorf(format, v...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	defLogger.logger.Sugar().Errorw(msg, keysAndValues...)
}

func Panic(msg string, fields ...Field) {
	defLogger.logger.Panic(msg, fields...)
}

func Panicf(format string, v ...interface{}) {
	defLogger.logger.Sugar().Panicf(format, v...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	defLogger.logger.Sugar().Panicw(msg, keysAndValues...)
}

func Fatal(msg string, fields ...Field) {
	defLogger.logger.Fatal(msg, fields...)
}

func Fatalf(format string, v ...interface{}) {
	defLogger.logger.Sugar().Fatalf(format, v...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	defLogger.logger.Sugar().Fatalw(msg, keysAndValues...)
}

func DebugSql(ctx string, t time.Time, sql string, v ...interface{}) {
	sql = strings.ReplaceAll(sql, "%", "%%")
	fmt.Sprintf("%s SQL: %s", ctx, sql)
	if len(v) >= 0 && len(v) < 50 {
		for i, val := range v {
			if val == nil || reflect.TypeOf(val).Kind() == reflect.String {
				sql = strings.Replace(sql, "?", "'%v'", 1)
			} else if reflect.TypeOf(val).Kind() == reflect.Slice {
				slice := fmt.Sprintf("{%v}", val)
				slice = strings.ReplaceAll(slice, " ", ",")
				slice = strings.ReplaceAll(slice, "[", "")
				slice = strings.ReplaceAll(slice, "]", "")
				v[i] = slice
				sql = strings.Replace(sql, "?", "'%v'", 1)
			} else if tm, ok := val.(time.Time); ok {
				v[i] = tm.Format("2006-01-02 15:04:05")
				sql = strings.Replace(sql, "?", "'%v'", 1)
			} else {
				sql = strings.Replace(sql, "?", "%v", 1)
			}
		}
		sql = fmt.Sprintf(sql, v...)
	} else {
		firstQuestionMark := strings.Index(sql, "?")
		sql = sql[:int(math.Min(float64(firstQuestionMark), float64(len(sql))))]
	}

	duration := time.Now().Sub(t).Seconds()
	Debugf("%s\nSQL execution time: %.4fs", sql, duration)
}

func WithContext(ctx context.Context) Logger {
	return defLogger.WithContext(ctx)
}

func AddContextHook(h ContextHook) {
	defLogger.AddContextHook(h)
}
