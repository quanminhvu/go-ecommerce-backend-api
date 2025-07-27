package initialize

import (
	"fmt"
	"time"

	"github.com/quanminhvu/go-ecommerce-backend-api/global"
	"github.com/quanminhvu/go-ecommerce-backend-api/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErr(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMySQL() {
	m := global.Config.MySQL
	dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, fmt.Sprintf("%d", m.Port), m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErr(err, "InitMySQL error")
	global.Mdb = db

	//set pool
	SetPool()
	migrateTable()

}

func SetPool() {
	m := global.Config.MySQL
	sqlDb, err := global.Mdb.DB()
	checkErr(err, "SetPool error")
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdle))
	sqlDb.SetMaxOpenConns(m.MaxOpen)
	sqlDb.SetConnMaxLifetime(time.Duration(m.LifeTime))
}

func migrateTable() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	checkErr(err, "migrateTable error")
}
