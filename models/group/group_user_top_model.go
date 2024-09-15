package models

import "server/common/models"

// GroupUserTopModel 用户置顶群聊表
type GroupUserTopModel struct {
	models.Model
	UserID  uint `json:"userID"`
	GroupID uint `json:"groupID"`
}
