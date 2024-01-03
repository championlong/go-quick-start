package app

import (
	"fmt"
	"testing"
)

func TestApp(t *testing.T) {
	NewApp("Mock Server",
		"mock-basename",
		WithDescription("mock描述信息"),
		WithRunFunc(run()),
	).Run()

}

func run() RunFunc {
	return func(basename string) error {
		fmt.Println("init mock app")
		return nil
	}
}
