package log

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/championlong/go-quick-start/pkg/log/utils"
	"github.com/natefinch/lumberjack"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLoggerAdapter struct {
	logger       *zap.Logger
	contextHooks ContextHooks
}

var zapConfig *ZapOptions

func New(config *ZapOptions) *ZapLoggerAdapter {
	if config == nil {
		config = NewZapOptions()
	}

	zapConfig = config
	if ok, _ := utils.PathExists(zapConfig.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", zapConfig.Director)
		_ = os.Mkdir(zapConfig.Director, os.ModePerm)
	}
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/server_debug.log", zapConfig.Director), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_info.log", zapConfig.Director), infoPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_warn.log", zapConfig.Director), warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_error.log", zapConfig.Director), errorPriority),
	}
	logger := zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if zapConfig.ShowLine {
		logger = logger.WithOptions(zap.AddCallerSkip(1), zap.AddCaller())
	}

	hostname, _ := os.Hostname()
	base := &baseInfo{
		hostname: hostname,
		program:  path.Base(os.Args[0]),
		ip:       utils.GetLocalIP(),
		pid:      os.Getpid(),
	}
	fields := []zapcore.Field{
		zap.Object(metaFieldName, base),
	}

	logger = logger.With(fields...)
	zap.ReplaceGlobals(logger)

	loggerInfo := &ZapLoggerAdapter{
		logger: logger,
	}

	return loggerInfo
}

// getEncoderConfig 获取zapcore.EncoderConfig.
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  zapConfig.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	switch {
	case zapConfig.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case zapConfig.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case zapConfig.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case zapConfig.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder.
func getEncoder() zapcore.Encoder {
	if zapConfig.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core.
func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := GetWriteSyncer(fileName) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(getEncoder(), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式.
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(zapConfig.Prefix + "2006/01/02 - 15:04:05.000"))
}

func GetWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // 日志文件的位置
		MaxSize:    10,   // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 200,  // 保留旧文件的最大个数
		MaxAge:     30,   // 保留旧文件的最大天数
		Compress:   true, // 是否压缩/归档旧文件
	}

	if zapConfig.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}

func (l *ZapLoggerAdapter) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, fields...)
}

func (l *ZapLoggerAdapter) Debugf(format string, v ...interface{}) {
	l.logger.Sugar().Debugf(format, v...)
}

func (l *ZapLoggerAdapter) Debugw(msg string, keysAndValues ...interface{}) {
	l.logger.Sugar().Debugw(msg, keysAndValues...)
}

func (l *ZapLoggerAdapter) Info(msg string, fields ...Field) {
	l.logger.Info(msg, fields...)
}

func (l *ZapLoggerAdapter) Infof(format string, v ...interface{}) {
	l.logger.Sugar().Infof(format, v...)
}

func (l *ZapLoggerAdapter) Infow(msg string, keysAndValues ...interface{}) {
	l.logger.Sugar().Infow(msg, keysAndValues...)
}

func (l *ZapLoggerAdapter) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, fields...)
}

func (l *ZapLoggerAdapter) Warnf(format string, v ...interface{}) {
	l.logger.Sugar().Warnf(format, v...)
}

func (l *ZapLoggerAdapter) Warnw(msg string, keysAndValues ...interface{}) {
	l.logger.Sugar().Warnw(msg, keysAndValues...)
}

func (l *ZapLoggerAdapter) Error(msg string, fields ...Field) {
	l.logger.Error(msg, fields...)
}

func (l *ZapLoggerAdapter) Errorf(format string, v ...interface{}) {
	l.logger.Sugar().Errorf(format, v...)
}

func (l *ZapLoggerAdapter) Errorw(msg string, keysAndValues ...interface{}) {
	l.logger.Sugar().Errorw(msg, keysAndValues...)
}

func (l *ZapLoggerAdapter) Panic(msg string, fields ...Field) {
	l.logger.Panic(msg, fields...)
}

func (l *ZapLoggerAdapter) Panicf(format string, v ...interface{}) {
	l.logger.Sugar().Panicf(format, v...)
}

func (l *ZapLoggerAdapter) Panicw(msg string, keysAndValues ...interface{}) {
	l.logger.Sugar().Panicw(msg, keysAndValues...)
}

func (l *ZapLoggerAdapter) Fatal(msg string, fields ...Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *ZapLoggerAdapter) Fatalf(format string, v ...interface{}) {
	l.logger.Sugar().Fatalf(format, v...)
}

func (l *ZapLoggerAdapter) Fatalw(msg string, keysAndValues ...interface{}) {
	l.logger.Sugar().Fatalw(msg, keysAndValues...)
}

func (l *ZapLoggerAdapter) AddContextHook(h ContextHook) {
	l.contextHooks = append(l.contextHooks, h)
}

func (l *ZapLoggerAdapter) WithContext(ctx context.Context) Logger {
	s := l.execContextHooks(ctx)
	if len(s) == 0 {
		return l
	}
	return &ZapLoggerAdapter{
		logger:       l.logger.With(s...),
		contextHooks: append(l.contextHooks[:0:0], l.contextHooks...),
	}
}

func (l *ZapLoggerAdapter) execContextHooks(ctx context.Context) []Field {
	var fields []Field
	if ctx != nil && len(l.contextHooks) != 0 {
		for _, h := range l.contextHooks {
			fields = append(fields, h(ctx)...)
		}
	}
	return fields
}
