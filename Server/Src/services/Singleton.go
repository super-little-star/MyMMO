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
		mlog.Info.Println("network package Singleton Init success...")
	})
}
func Instance() *singleton {
	if instance != nil {
		return instance
	}
	return nil
}
