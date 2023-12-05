package app

import (
	"github.com/championlong/go-quick-start/internal/pkg/app"
)

func NewApp(basename string) *app.App {
	application := app.NewApp("Ging API Server",
		basename,
		app.WithDescription("Web脚手架"),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run()),
	)
	return application
}

func run() app.RunFunc {
	return func(basename string) error {
		return Run()
	}
}
