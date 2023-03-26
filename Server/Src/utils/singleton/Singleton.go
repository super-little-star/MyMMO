package singleton

import (
	"mmo_server/network"
	"mmo_server/utils/mlog"
	"sync"
)

type singleton struct {
	MessageHandleCenter *network.GMessageHandleCenter
}

var once sync.Once
var instance *singleton

func InitSingleton() {
	once.Do(func() {
		instance = &singleton{
			MessageHandleCenter: &network.GMessageHandleCenter{},
		}
		mlog.Info.Println("Singleton Init success...")
	})
}
func Instance() *singleton {
	if instance != nil {
		return instance
	}
	return nil
}
