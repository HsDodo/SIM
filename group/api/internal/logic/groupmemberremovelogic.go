package logic

import (
	"context"
	"errors"
	group "server/models/group"

	"server/group/api/internal/svc"
	"server/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMemberRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMemberRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMemberRemoveLogic {
	return &GroupMemberRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupMemberRemoveLogic) GroupMemberRemove(req *types.GroupMemberRemoveRequest) (resp *types.NoResponse, err error) {
	// Note: 退群或者踢人
	// Note: 谁能调这个接口 必须得是这个群的成员
	var member group.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", req.GroupID, req.UserID).Error
	if err != nil {
		return nil, errors.New("违规调用")
	}

	var groupModel group.GroupModel
	err = l.svcCtx.DB.Take(&groupModel, req.GroupID).Error
	if err != nil {
		return nil, errors.New("群不存在")
	}

	// 用户自己退群
	if req.UserID == req.MemberID {
		if member.Role == 2 {
			return nil, errors.New("群主不能退群，只能解散群聊")
		}
		l.svcCtx.DB.Delete(&member)
		// 给群验证表里面加条记录
		l.svcCtx.DB.Create(&group.GroupVerifyModel{
			GroupID: member.GroupID,
			UserID:  req.UserID,
			Type:    2, // 退群
		})
		return
	}
	// 把用户踢出群聊
	var kickOutMember group.GroupMemberModel
	err = l.svcCtx.DB.Preload("MsgList").Take(&kickOutMember, "group_id = ? and user_id = ?", req.GroupID, req.MemberID).Error
	if err != nil {
		return nil, errors.New("该用户不是群成员")
	}
	if (member.Role == 0) || (member.Role == 2 && kickOutMember.Role == 2) || (member.Role == 1 && kickOutMember.Role == 1) || (member.Role == 1 && kickOutMember.Role == 2) {
		return nil, errors.New("角色权限错误")
	}
	//在删之前解除这个用户的关联消息，也就是在群里能看到消息，但是不能通过消息访问这个人
	if len(kickOutMember.MsgList) > 0 {
		err := l.svcCtx.DB.Model(&kickOutMember.MsgList).Update("group_member_id", nil).Error
		if err != nil {
			return nil, errors.New("解除用户关联消息失败")
		}
		logx.Infof("解除用户关联消息%d条", len(kickOutMember.MsgList))
	}
	err = l.svcCtx.DB.Delete(&kickOutMember).Error
	if err != nil {
		logx.Errorf("移除成员失败：%v", err)
		return nil, errors.New("移除成员失败")
	}

	return
}
