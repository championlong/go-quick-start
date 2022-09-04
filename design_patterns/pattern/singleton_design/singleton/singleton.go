package singleton

import "sync"

// Singleton 饿汉式单例
type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func GetInstance() *Singleton {
	return singleton
}

// Singleton 懒汉式单例
var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

// GetLazyInstance 懒汉式(双重检测)
func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}