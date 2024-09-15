package model

import "gorm.io/gorm"

type LogModel struct {
	gorm.Model
	LogType      int8   `json:"logType"` // 日志类型  1 操作日志 2 运行日志
	IP           string `gorm:"size:32" json:"ip"`
	Addr         string `gorm:"size:64" json:"addr"`
	UserID       uint   `json:"userID"`
	UserNickname string `gorm:"size:64" json:"userNickname"`
	UserAvatar   string `gorm:"size:256" json:"userAvatar"`
	Level        string `gorm:"size:12" json:"level"`
	Title        string `gorm:"size:32" json:"title"`
	Content      string `json:"content"`                // 日志详情
	Service      string `gorm:"size:32" json:"service"` // 服务  记录微服务的名称
	IsRead       bool   `json:"isRead"`
}
