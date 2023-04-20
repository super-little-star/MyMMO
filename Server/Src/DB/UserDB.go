package DB

import (
	"database/sql"
	"errors"
	"mmo_server/DB/Model"
	"mmo_server/utils/mlog"
)

var (
	ErrUserNameExist = errors.New("DB :: UserName is Exist") // 错误：用户名已存在
	ErrUserNotExist  = errors.New("DB :: User is not Exist") // 错误：用户不存在
)

// Register
//
//	@Description: DB的注册操作
//	@param uid 用户ID
//	@param userName 用户名
//	@param psw 用户密码
//	@param rt 注册时间
//	@return error 错误
func Register(uid int64, userName string, psw string, rt int64) error {
	if _, err := dB.Exec("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE"); err != nil {
		return err
	}
	tx, err := dB.Begin()

	defer func() {
		if err := tx.Rollback(); err != nil {
			mlog.Error.Println("SQL Rollback err:%v", err)
		}
	}()
	if err != nil {
		mlog.Error.Println("Transaction begin is error: %v", err)
	}

	s := "SELECT userName FROM DBUser WHERE userName = ? LIMIT 1"
	row := tx.QueryRow(s, userName)

	var name string
	err = row.Scan(&name)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			{
				// 查询为空，则插入数据
				i := "INSERT INTO DBUser (uid,userName,password,registerTime) VALUES (?,?,?,?)"
				_, err = tx.Exec(i, uid, userName, psw, rt)
				if err != nil {
					return err
				}

				if err := tx.Commit(); err != nil {
					mlog.Error.Println("SQL Commit error:%v", err)
				}
				return nil
			}
		default:
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		mlog.Error.Println("SQL Commit error:%v", err)
	}
	return ErrUserNameExist
}

func GetDbUser(userName string) (*Model.DbUser, error) {
	s := "SELECT UID,Password,CharacterCount,RegisterTime FROM DBUser WHERE userName = ? LIMIT 1"
	row := dB.QueryRow(s, userName)
	user := &Model.DbUser{}
	if err := row.Scan(&user.UID, &user.Password, &user.CharacterCount, &user.RegisterTime); err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrUserNotExist
		default:
			return nil, err
		}
	}
	user.UserName = userName
	return user, nil
}
