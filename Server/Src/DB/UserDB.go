package DB

import (
	"database/sql"
	"errors"
)

var (
	ErrUserNameExist = errors.New("DB :: UserName is Exist") // 错误：用户名已存在
)

func UserRegister(uid int64, userName string, psw string, rt int64) error {
	s := "SELECT userName FROM DBUser WHERE userName = ? LIMIT 1"
	row := dB.QueryRow(s, userName)
	var name string
	err := row.Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows { // 判断查询结果是否为空
			// 查询为空，则插入数据
			i := "INSERT INTO DBUser (uid,userName,password,registerTime) VALUES (?,?,?,?)"
			_, err = dB.Exec(i, uid, userName, psw, rt)
			if err != nil {
				return ErrSQL
			}
			return nil
		} else {
			return ErrSQL
		}

	}

	return ErrUserNameExist
}
