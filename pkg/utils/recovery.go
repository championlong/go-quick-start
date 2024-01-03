package utils

import (
	"github.com/championlong/go-quick-start/pkg/log"
	"runtime"
)

func Recovery() {
	if r := recover(); r != nil {
		buf := make([]byte, 1<<18)
		n := runtime.Stack(buf, false)
		log.Errorf(" %+v. stack: %s", r, buf[0:n])
	}
}

type RecoveryFallBackFunc func(interface{})
