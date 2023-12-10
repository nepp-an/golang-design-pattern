package Singleton

import "sync"

type Singleton struct{}

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func GetSingleInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}