package app

import (
	"fmt"
	"testing"
)

type mockService struct {
}

func (m mockService) Start() error {
	fmt.Println("start")
	return nil
}

func (m mockService) Stop() error {
	fmt.Println("end")
	return nil
}

func TestService(t *testing.T) {
	RunService(&mockService{}).Wait()
}
