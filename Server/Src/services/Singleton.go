package services

import (
	"mmo_server/utils/mlog"
	"sync"
)

type singleton struct {
	UserService *GUserService
}

var once sync.Once
var instance *singleton

func InitSingleton() {
	once.Do(func() {
		instance = &singleton{
			UserService: &GUserService{},
		}
		mlog.Info.Println("Service package Singleton Init success...")
	})
}
func Instance() *singleton {
	if instance != nil {
		return instance
	}
	mlog.Error.Println("Instance is Null!!!")
	return nil
}
