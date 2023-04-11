package DB

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dB *sql.DB

var IDGenerator *GIDGenerator

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

	if g, err := NewIDGenerator(1); err != nil {
		return err
	} else {
		IDGenerator = g
	}

	return nil
}

// GetDB
//
//	@Description: 获取DB实体
//	@return *sql.DB
func GetDB() *sql.DB {
	if dB != nil {
		return dB
	} else {
		log.Println("DB is not Init!!!")
		return nil
	}
}

func GetIDGenerator() *GIDGenerator {
	if IDGenerator != nil {
		return IDGenerator
	} else {
		log.Println("ID Generator is not Init!!!")
		return nil
	}
}

func Stop() {
	if err := dB.Close(); err != nil {

		return
	}
	dB = nil
}
