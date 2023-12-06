package utils

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/braintree/manners"
	"github.com/gorilla/mux"
)

func StartServer() {
	gracefulServer := manners.NewWithServer(&http.Server{
		Addr:         "8080",
		IdleTimeout:  1 * time.Minute, //长连接维持多久
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGPIPE)
	go func() {
		for {
			sig := <-c
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				gracefulServer.Close()
				return
			case syscall.SIGPIPE:
				// ignore
			}
		}
	}()

	r := http.NewServeMux()
	r.Handle("/", routers())
	gracefulServer.Handler = r

	// 启动 HTTP server
	err := gracefulServer.ListenAndServe()
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func routers() *mux.Router {
	r := mux.NewRouter()
	return r
}
