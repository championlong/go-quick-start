package main

import (
	"github.com/championlong/backend-common/demo/gin/router"
)

func main() {
	r := router.Routers()
	r.Run(":8080")
}
