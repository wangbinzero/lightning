package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	//定义全局准备数据库
	MainDb, BackupDb *gorm.DB
)

type GormDB struct {
	*gorm.DB
	gdbDone bool
}
