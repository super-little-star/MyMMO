package manager

import (
	"errors"
	"mmo_server/DB"
	"mmo_server/DB/Model"
	"mmo_server/utils/gencrypt"
	"mmo_server/utils/xuid"
	"time"
)

var ErrPasswordNotMatch = errors.New("manager:: user password is not match")

type IUserManager interface {
	UserRegister(userName string, psw string) error
	UserLogin(userName string, psw string) (*Model.DbUser, error)
	CreateCharacter(uid int64, name string, class int) ([]*Model.DbCharacter, error)
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

	newPsw, err := gencrypt.EncryptPassword(psw)

	if err != nil {
		return err
	}

	if err := DB.Register(uid, userName, newPsw, time.Now().Unix()); err != nil {
		return err
	}
	return nil
}

// UserLogin
//
//	@Description: 用户登录处理
//	@receiver u
//	@param userName 用户名
//	@param psw 密码
//	@return *Model.DbUser 返回一个用户实体
//	@return error
func (u *GUserManager) UserLogin(userName string, psw string) (*Model.DbUser, error) {
	user, err := DB.GetDbUser(userName)
	if err != nil {
		return nil, err
	}
	// 密码校验
	if !gencrypt.EqualsPassword(psw, user.Password) {
		return nil, ErrPasswordNotMatch
	}

	if user.Characters, err = DB.GetCharacters(user.UID); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *GUserManager) CreateCharacter(uid int64, name string, class int) ([]*Model.DbCharacter, error) {
	if err := DB.CreateCharacter(uid, name, class, time.Now().Unix()); err != nil {
		return nil, err
	}

	return DB.GetCharacters(uid)
}
