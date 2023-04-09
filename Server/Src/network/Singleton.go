package network

import (
	"mmo_server/utils/mlog"
	"sync"
)

type singleton struct {
	MessageHandleCenter IMessageHandleCenter
	MessageHandOut      IMessageHandOut
	ConnectionManager   IConnectionManager
}

var once sync.Once
var instance *singleton

func InitSingleton() {
	once.Do(func() {
		instance = &singleton{
			MessageHandleCenter: &GMessageHandleCenter{},
			MessageHandOut:      &GMessageHandOut{},
			ConnectionManager:   &GConnectionManager{},
		}
		instance.MessageHandleCenter.Init()
		instance.MessageHandOut.Init()
		instance.ConnectionManager.Init()
		mlog.Info.Println("network package Singleton Init success...")
	})
}
func Instance() *singleton {
	if instance != nil {
		return instance
	}
	return nil
}
