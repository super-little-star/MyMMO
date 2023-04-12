package manager

import (
	"mmo_server/DB"
	"mmo_server/utils/encrypt"
	"mmo_server/utils/xuid"
	"time"
)

type IUserManager interface {
	UserRegister(userName string, psw string) error
}

type GUserManager struct {
}

func NewUserManager() IUserManager {
	return &GUserManager{}
}

// UserRegister
//
//	@Description: 处理用户注册
//	@receiver u
//	@param userName
//	@param psw
//	@return error
func (u *GUserManager) UserRegister(userName string, psw string) error {
	uid := xuid.Generator.NextId()

	newPsw, err := encrypt.EncryptPassword(psw)

	if err != nil {
		return err
	}

	if err := DB.UserRegister(uid, userName, newPsw, time.Now().Unix()); err != nil {
		return err
	}

	return nil
}
