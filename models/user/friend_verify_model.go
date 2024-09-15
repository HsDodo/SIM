package models

import (
	"gorm.io/gorm"
	"server/common/models"
)

type FriendVerifyModel struct {
	gorm.Model
	SendUserID           uint                         `gorm:"not null" json:"sendUserID"`           // 发起验证方ID
	RevUserID            uint                         `gorm:"not null" json:"revUserID"`            // 接受验证方ID
	Status               int8                         `gorm:"default:0" json:"status"`              // 状态 0-未操作 1-同意 2-拒绝 3-忽略 4-删除
	SendStatus           int8                         `gorm:"default:0" json:"sendStatus"`          // 发送方状态 4 删除
	RevStatus            int8                         `gorm:"default:0" json:"revStatus"`           // 接收方状态 0 未操作 1 同意 2 拒绝 3 忽略 4 删除
	AdditionalMessages   string                       `gorm:"size:128" json:"additionalMessages"`   // 附加消息
	VerificationQuestion *models.VerificationQuestion `gorm:"size:128" json:"verificationQuestion"` // 验证问题  为3和4的时候需要
	SendTime             int64                        `gorm:"autoCreateTime" json:"sendTime"`       // 发送时间
	// 关联查询
	SendUser UserModel `gorm:"foreignKey:SendUserID" json:"-"` // 发起验证方 （关联查询用）
	RevUser  UserModel `gorm:"foreignKey:RevUserID" json:"-"`  // 接受验证方	（关联查询用）
}
