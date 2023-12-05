package initialize

import (
	"fmt"
	"github.com/championlong/go-quick-start/internal/app/global"
	"github.com/championlong/go-quick-start/internal/pkg/options"
	"github.com/championlong/go-quick-start/internal/pkg/utils"
)

func Timer() {
	if global.GVA_CONFIG.Timer.Start {
		for i := range global.GVA_CONFIG.Timer.Detail {
			go func(detail options.Detail) {
				global.GVA_Timer.AddTaskByFunc("ClearDB", global.GVA_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.GVA_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(global.GVA_CONFIG.Timer.Detail[i])
		}
	}
}
