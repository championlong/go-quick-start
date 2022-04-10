package main

import (
	"context"
	"fmt"
	"github.com/championlong/backend-common/app/global"
	"github.com/championlong/backend-common/app/initialize"
	"github.com/go-redis/redis/v8"
	juju_ratelimit "github.com/juju/ratelimit"
	"time"
)

func setup() {
	initialize.Redis()
}

var roidUserLimiter = juju_ratelimit.NewBucketWithQuantum(time.Second, 500, 500)

func main() {
	setup()
	go dealDevice()
	select {}
}

//扫描所有device
func dealDevice() {

	pipelineRta := global.GVA_REDIS.Pipeline()
	var cursor uint64
	var n int
	start := time.Now()
	for {
		var keys []string
		var err error
		delEncrypt := make([]*redis.IntCmd, 0)
		roidUserLimiter.Wait(1)
		keys, cursor, err = global.GVA_REDIS.Scan(context.Background(),cursor, "rud:*", 500).Result()
		if err != nil {
			fmt.Println("scan error ", err.Error())
		}
		n += len(keys)
		for index := range keys {
			delEncrypt = append(delEncrypt, pipelineRta.Del(context.Background(),keys[index]))
		}
		_, _ = pipelineRta.Exec(context.Background())
		if cursor == 0 {
			break
		}
	}
	fmt.Printf("found %d keys\n", n)
	fmt.Println("已经扫描完毕")
	fmt.Println(time.Now().Sub(start))
}
