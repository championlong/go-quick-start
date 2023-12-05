package log

import (
	"context"
	"go.uber.org/zap/zapcore"
)

const (
	consoleFormat = "console"
	jsonFormat    = "json"

	metaFieldName = "_base_"

	defaultDirector = "server_log"
)

// Field is an alias for the field structure in the underlying log frame.
type Field = zapcore.Field

// Level is an alias for the level structure in the underlying log frame.
type Level = zapcore.Level

type ContextHook func(context.Context) []Field

type ContextHooks []ContextHook

type baseInfo struct {
	hostname string
	program  string
	ip       string
	pid      int
}

func (m *baseInfo) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if m.hostname != "" {
		enc.AddString("hostname", m.hostname)
	}
	if m.program != "" {
		enc.AddString("program", m.program)
	}
	if m.ip != "" {
		enc.AddString("ip", m.ip)
	}
	enc.AddInt("pid", m.pid)
	return nil
}

type ZapOptions struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // 输出
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                  // 日志文件夹
	ShowLine      bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 显示行
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 输出控制台
}

// NewZapOptions creates an Options object with default parameters.
func NewZapOptions() *ZapOptions {
	return &ZapOptions{
		Director:      defaultDirector,
		ShowLine:      true,
		Format:        jsonFormat,
		StacktraceKey: "stacktrace",
		LogInConsole:  true,
	}
}
