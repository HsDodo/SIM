package models

import (
	"server/common/models"
	"time"
)

type GroupMsgModel struct {
	models.Model
	GroupID          uint              `json:"group_id"`                          // 群组ID
	SenderID         uint              `json:"sender_id"`                         // 发送者ID
	GroupMemberID    uint              `json:"groupMemberID"`                     // 群成员id
	GroupMemberModel *GroupMemberModel `gorm:"foreignKey:GroupMemberID" json:"-"` // 对应的群成员
	MsgType          models.MsgType    `json:"msg_type"`                          // 消息类型 0-文本 1-图片 2-文件 3-音频 4-视频 5-语音通话 6-视频通话 7-撤回消息 8-转发消息 9-回复消息 10-@消息
	MsgPreview       string            `gorm:"size:64" json:"msg_preview"`        // 消息预览
	Msg              models.Message    `json:"msg"`                               // 消息内容
	SystemMsg        *models.SystemMsg `json:"system_msg"`
	Timestamp        time.Time         `json:"timestamp"` // 消息时间
}

func (chat GroupMsgModel) MsgPreviewMethod() string {
	if chat.SystemMsg != nil {
		switch chat.SystemMsg.Type {
		case 1:
			return "[系统消息]- 该消息涉黄，已被系统拦截"
		case 2:
			return "[系统消息]- 该消息涉恐，已被系统拦截"
		case 3:
			return "[系统消息]- 该消息涉政，已被系统拦截"
		case 4:
			return "[系统消息]- 该消息不正当言论，已被系统拦截"
		}
		return "[系统消息]"
	}
	return chat.Msg.MsgPreview()
}
