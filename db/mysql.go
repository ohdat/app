package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/zcpua/manager"
	"log"
	"sync"
	"time"
)

var mysqlDB *sql.DB

func initMysql() {
	var err error
	host := viper.GetString("mysql.host")
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	database := viper.GetString("mysql.database")
	port := viper.GetInt("mysql.port")
	mysqlDB, err =
		manager.New(database, user, password, host).Set(
			manager.SetCharset("utf8mb4"),
			manager.SetAllowCleartextPasswords(true),
			manager.SetInterpolateParams(true),
			manager.SetTimeout(10*time.Second),
			manager.SetReadTimeout(10*time.Second),
			manager.SetParseTime(true),
		).Port(port).Open(true)
	if err != nil {
		log.Fatalln("mysql Init Error: ", err)
	}
	fmt.Println("mysql Init finished")
}

var mysqlOnce sync.Once

func GetMysql() *sql.DB {
	mysqlOnce.Do(func() {
		initMysql()
	})
	return mysqlDB
}
