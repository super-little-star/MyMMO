package DB

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB

// Init
//
//	@Description: 初始化数据库
//	@param ip 数据库IP地址
//	@param dbName 名称
//	@return error
func Init(user string, psw string, ip string, port string, dbName string) error {
	if dB != nil {
		if err := dB.Close(); err != nil {
			return err
		}
		dB = nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, psw, ip, port, dbName)
	if db, err := sql.Open("mysql", dsn); err != nil {
		return err
	} else {
		dB = db
	}

	if err := dB.Ping(); err != nil {
		return err
	}
	return nil
}

func Stop() {
	if err := dB.Close(); err != nil {

		return
	}
	dB = nil
}
