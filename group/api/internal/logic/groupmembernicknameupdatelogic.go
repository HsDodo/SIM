package logic

import (
	"context"
	"errors"
	group "server/models/group"

	"server/group/api/internal/svc"
	"server/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMemberNicknameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMemberNicknameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMemberNicknameUpdateLogic {
	return &GroupMemberNicknameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupMemberNicknameUpdateLogic) GroupMemberNicknameUpdate(req *types.GroupMemberNicknameUpdateRequest) (resp *types.NoResponse, err error) {
	// todo: add your logic here and delete this line
	var member group.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", req.GroupID, req.UserID).Error
	if err != nil {
		return nil, errors.New("违规调用")
	}
	var operatedMember group.GroupMemberModel
	err = l.svcCtx.DB.Take(&operatedMember, "group_id = ? and user_id = ?", req.GroupID, req.MemberID).Error
	if err != nil {
		return nil, errors.New("该用户不是群成员")
	}
	// 自己修改自己的
	if req.UserID == req.MemberID {
		l.svcCtx.DB.Model(&member).Updates(map[string]interface{}{
			"nickname": req.Nickname,
		})
		return
	}
	// 修改别人的，只有管理员和群主可以修改, 管理员不能修改群主的
	// 角色 0 普通成员 1 管理员 2 群主
	if (member.Role == 0 && (operatedMember.Role == 1 || operatedMember.Role == 2)) || (member.Role == 1 && operatedMember.Role == 2) {
		return nil, errors.New("用户角色权限 错误")
	}

	err = l.svcCtx.DB.Model(&operatedMember).Updates(map[string]interface{}{
		"nickname": req.Nickname,
	}).Error
	if err != nil {
		logx.Errorf("修改用户昵称失败: %v", err)
		return nil, errors.New("修改用户昵称失败")
	}
	return
}
