package db

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var gormDB *gorm.DB

func initGorm() *gorm.DB {
	var sqlDB = GetMysql()
	ormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		log.Fatalln("gorm Init Error: ", err)
	}
	return ormDB
}

var gormOnce sync.Once

func GetGorm() *gorm.DB {
	gormOnce.Do(func() {
		gormDB = initGorm()
	})
	return gormDB
}
