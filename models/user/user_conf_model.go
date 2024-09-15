package models

import (
	models "server/common/models"
)

type UserConfModel struct {
	models.Model
	UserID               uint                         `json:"userId"`               // 用户ID
	RecallMessage        *string                      `json:"recallMessage"`        // 撤回消息的提示内容
	FriendOnline         bool                         `json:"friendOnline"`         // 好友上线提醒
	Sound                bool                         `json:"sound"`                // 声音提醒
	SecureLink           bool                         `json:"secureLink"`           // 安全链接
	SavePwd              bool                         `json:"savePwd"`              // 保存密码
	SearchUser           int8                         `json:"searchUser"`           // 搜索用户 0-不允许查找 1-用户号查找 2-昵称查找
	Verification         int8                         `json:"friendVerification"`   // 好友验证 0-不允许添加 1-允许任何人添加 2-验证消息 3-需要回答问题 4-需要正确回答问题
	VerificationQuestion *models.VerificationQuestion `json:"verificationQuestion"` // 好友验证问题, 当FriendVerification为3或4时有效 ，1对多关系
	OnlineStatus         bool                         `json:"onlineStatus"`         // 在线状态 0-不在线 1-在线
	ForbidChat           bool                         `json:"fobidChat"`            // 限制聊天
	ForbidAddUser        bool                         `json:"fobidAddUser"`         // 限制加人
	ForbidCreateGroup    bool                         `json:"fobidCreateGroup"`     // 限制建群
	ForbidInGroupChat    bool                         `json:"fobidInGroupChat"`     // 限制加群
}

// ProblemCount 问题的个数
func (uc UserConfModel) ProblemCount() (c int) {
	if uc.VerificationQuestion != nil {
		if uc.VerificationQuestion.Problem1 != nil {
			c += 1
		}
		if uc.VerificationQuestion.Problem2 != nil {
			c += 1
		}
		if uc.VerificationQuestion.Problem3 != nil {
			c += 1
		}
	}
	return
}
