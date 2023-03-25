package singleton

import (
	"mmo_server/network"
	"sync"
)

type singleton struct {
	MessageHandleCenter *network.MessageHandleCenter
}

var once sync.Once
var instance *singleton

func InitSingleton() {
	once.Do(func() {
		instance = &singleton{
			MessageHandleCenter: &network.MessageHandleCenter{},
		}
	})
}
func GetInstance() *singleton {
	if instance != nil {
		return instance
	}
	return nil
}
