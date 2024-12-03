package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func init() {
	InitMysql()
}

// InitMysql 初始化mysql
func InitMysql() (*gorm.DB, error) {
	dsn := "ats:EVVnrLA9kCjY5aNuwInqujpEnM6GyrdY@tcp(10.150.19.18:3318)/ats?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	dbMysql, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println("open db_mysql error ", err)
		return nil, err
	}

	// 是否打开日志
	dbMysql.Debug()
	DB = dbMysql

	//创表
	// dbMysql.AutoMigrate()
	//db, error := dbMysql.DB()
	////设置连接池的最大闲置连接数
	//db.set
	//db.SetMaxIdleConns(10)
	////设置连接池中的最大连接数量
	//db.SetMaxOpenConns(100)
	////设置连接的最大复用时间
	//db.SetConnMaxLifetime(10 * time.Second)
	return dbMysql, nil
}
