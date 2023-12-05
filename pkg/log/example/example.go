package main

import (
	"context"
	"github.com/championlong/go-quick-start/pkg/log"
)

func main() {
	log.AddContextHook(ServiceContextHook)
	log.WithContext(context.Background()).Debugf("1111")
	log.Debugf("1111")
}

func ServiceContextHook(ctx context.Context) []log.Field {
	var fields []log.Field
	if ctx != nil {
		fields = append(fields, log.String("test", "测试信息"))
	}
	return fields
}
