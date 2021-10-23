package core

import (
	"sync"
)

var once sync.Once

var instance *Router

func GetInstanceRouter() *Router {
	once.Do(func(){
		instance = New()
	})

	return instance
}
