package models

import (
	"fmt"
	"gin-blog/pkg/setting"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

var (
	err                                               error
	dbType, dbName, user, password, host, tablePrefix string
)

func init() {

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	Db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName,
	))
	if err != nil {
		log.Fatalf("Fail to open db err : %v", err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName

	}

	Db.SingularTable(true) //创建表默认在表名后加s,全局设置表名不可以为复数形式。
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)

}

func CloseDB() {
	defer Db.Close()
}
