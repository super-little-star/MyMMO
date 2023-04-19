package DB

import (
	"database/sql"
	"errors"
	"mmo_server/utils/mlog"
)

var (
	ErrUserNameExist = errors.New("DB :: UserName is Exist") // 错误：用户名已存在
)

func UserRegister(uid int64, userName string, psw string, rt int64) error {
	if _, err := dB.Exec("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE"); err != nil {
		return ErrSQL
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
		if err == sql.ErrNoRows { // 判断查询结果是否为空
			// 查询为空，则插入数据
			i := "INSERT INTO DBUser (uid,userName,password,registerTime) VALUES (?,?,?,?)"
			_, err = tx.Exec(i, uid, userName, psw, rt)
			if err != nil {
				return ErrSQL
			}

			if err := tx.Commit(); err != nil {
				mlog.Error.Println("SQL Commit error:%v", err)
			}
			return nil
		} else {
			return ErrSQL
		}

	}
	if err := tx.Commit(); err != nil {
		mlog.Error.Println("SQL Commit error:%v", err)
	}
	return ErrUserNameExist
}
