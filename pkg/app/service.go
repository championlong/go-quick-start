package app

import (
	"github.com/championlong/go-quick-start/pkg/log"
	"github.com/championlong/go-quick-start/pkg/utils"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
)

type Service interface {
	Start() error
	Stop() error
}

type serviceRunner struct {
	signals chan os.Signal
	service Service

	stopped int32

	wg sync.WaitGroup
}

func RunService(s Service) *serviceRunner {
	r := newServiceRunner(s)
	r.run()
	return r
}

func newServiceRunner(s Service) *serviceRunner {
	return &serviceRunner{
		signals: make(chan os.Signal, 1),
		service: s,
	}
}

func (r *serviceRunner) run() {
	r.wg.Add(1)
	go r.handleSignal()
	go r.handleStart()
}

func (r *serviceRunner) handleStart() {
	func() {
		defer utils.Recovery()
		err := r.service.Start()
		if err != nil {
			log.Errorf("handler start: %s", err.Error())
		}
	}()
	if atomic.LoadInt32(&r.stopped) == 0 {
		r.wg.Done()
	}
}

func (r *serviceRunner) handleSignal() {
	signal.Notify(r.signals, syscall.SIGPIPE, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)
	for {
		select {
		case sig := <-r.signals:
			log.Infof("received signal: %s", sig)
			switch sig {
			case syscall.SIGPIPE:
			case syscall.SIGINT:
				r.handlerClose()
				log.Infof("Failure exit for systemd restarting")
				os.Exit(1)
			default:
				r.handlerClose()
				r.wg.Done()
			}
		}
	}
}

func (r *serviceRunner) handlerClose() {
	atomic.StoreInt32(&r.stopped, 1)
	err := r.service.Stop()
	if err != nil {
		log.Errorf("handler close: %s", err.Error())
	}
}

func (r *serviceRunner) Wait() {
	r.wg.Wait()
	log.Flush()
}
