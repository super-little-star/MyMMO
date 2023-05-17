package DB

import (
	"database/sql"
	"errors"
	"fmt"
	"mmo_server/DB/DbObject"
)

var (
	ErrUserNameExist      = errors.New("DB :: UserName is Exist")       // 错误：用户名已存在
	ErrUserNotExist       = errors.New("DB :: User is not Exist")       // 错误：用户不存在
	ErrCharacterNameExist = errors.New("DB :: Character Name is Exist") // 错误：角色名已存在
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
		_ = tx.Rollback()
	}()
	if err != nil {
		return err
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
					return err
				}
				return nil
			}
		default:
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return ErrUserNameExist
}

// GetDbUser
//
//	@Description: 获取用户信息
//	@param userName
//	@return *DbObject.DbUser
//	@return error
func GetDbUser(userName string) (*DbObject.DbUser, error) {
	s := "SELECT UID,Password,RegisterTime FROM DBUser WHERE userName = ? LIMIT 1"
	row := dB.QueryRow(s, userName)
	user := &DbObject.DbUser{}
	if err := row.Scan(&user.UID, &user.Password, &user.RegisterTime); err != nil {
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

// GetCharacters
//
//	@Description: 获取角色列表
//	@param user
//	@return error
func GetCharacters(uid int64) ([]*DbObject.DbCharacter, error) {

	var characters []*DbObject.DbCharacter

	str := "SELECT ID,UserID,Name,Class,Level From DbCharacter WHERE UserID = ?"
	rows, err := dB.Query(str, uid)
	defer func() {
		if err := rows.Close(); err != nil {
			fmt.Printf("rows close is error:%s", err)
		}
	}()

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return characters, nil
		default:
			return nil, err
		}
	}

	for rows.Next() {
		c := &DbObject.DbCharacter{}
		if err := rows.Scan(&c.ID, &c.UserID, &c.Name, &c.Class, &c.Level); err != nil {
			return nil, err
		}
		characters = append(characters, c)
	}

	return characters, nil
}

func CreateCharacter(uid int64, name string, class int, createTime int64) error {
	if _, err := dB.Exec("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE"); err != nil {
		return err
	}
	tx, err := dB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	s := "SELECT name FROM DbCharacter WHERE name = ? LIMIT 1"
	row := tx.QueryRow(s, name)
	var temp string
	err = row.Scan(&temp)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			{
				i := "INSERT INTO DbCharacter(userId,Name,Class,CreateTime) VALUES (?,?,?,?)"
				_, err = tx.Exec(i, uid, name, class, createTime)
				if err != nil {
					return err
				}
				if err := tx.Commit(); err != nil {
					return err
				}
				return nil
			}
		default:
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return ErrCharacterNameExist
}

func DeleteCharacter(uid int64, characterId int32) error {

	tx, err := dB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	s := "DELETE FROM DbCharacter WHERE ID = ? AND UserID = ?"

	ret, _ := tx.Exec(s, characterId, uid)

	ex, _ := ret.RowsAffected()

	if ex > 0 {
		_ = tx.Commit()
		return nil
	} else {
		return ErrSql
	}

}
