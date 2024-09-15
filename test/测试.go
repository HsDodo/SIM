package main

import (
	log "server/common/logger"
	"server/core"
)

func main() {

	gorm := core.InitGorm("root:root@tcp(127.0.0.1:3306)/sim_db?charset=utf8&parseTime=True&loc=Local")
	if gorm == nil {
		log.Error("gorm初始化失败")
		return
	}
	log.Info("gorm初始化成功")
}
