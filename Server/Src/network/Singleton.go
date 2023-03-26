package network

import (
	"mmo_server/utils/mlog"
	"sync"
)

type singleton struct {
	MessageHandleCenter *GMessageHandleCenter
	MessageHandOut      *GMessageHandOut
}

var once sync.Once
var instance *singleton

func InitSingleton() {
	once.Do(func() {
		instance = &singleton{
			MessageHandleCenter: &GMessageHandleCenter{},
			MessageHandOut:      &GMessageHandOut{},
		}
		mlog.Info.Println("network package Singleton Init success...")
	})
}
func Instance() *singleton {
	if instance != nil {
		return instance
	}
	return nil
}
