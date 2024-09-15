package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	log "server/common/logger"
	"time"
)

func InitGorm(dsn string) *gorm.DB {
	var mylogger logger.Interface
	mylogger = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mylogger,
	}) //db 是 *gorm.DB
	if err != nil {
		log.Fatalf("gorm初始化失败!: %v", err)
		return nil
	}
	sqlDB, _ := db.DB()                     //db.DB() 获取底层sql.DB
	sqlDB.SetMaxIdleConns(10)               // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)              // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大存活时间
	return db
}
