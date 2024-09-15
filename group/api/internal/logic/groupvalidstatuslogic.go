package logic

import (
	"context"
	"errors"
	models "server/models/group"

	"server/group/api/internal/svc"
	"server/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupValidStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupValidStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupValidStatusLogic {
	return &GroupValidStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupValidStatusLogic) GroupValidStatus(req *types.GroupValidStatusRequest) (resp *types.NoResponse, err error) {
	// Note: 群验证状态更新
	var groupValidModel models.GroupVerifyModel
	err = l.svcCtx.DB.Take(&groupValidModel, req.ValidID).Error
	if err != nil {
		return nil, errors.New("不存在的验证记录")
	}
	switch req.Status {
	case 1, 2, 3:
		if groupValidModel.Status != 0 {
			return nil, errors.New("已经处理过该验证请求了")
		}
	case 4:
		if groupValidModel.Status == 0 {
			return nil, errors.New("只能删除处理过的请求")
		}
	default:
		return nil, errors.New("错误的状态")
	}

	// 判断我有没有权限处理这个请求
	var member models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "user_id = ? and group_id = ?", req.UserID, groupValidModel.GroupID).Error
	if err != nil {
		return nil, errors.New("没有处理该操作的权限")
	}

	if member.Role == 0 {
		return nil, errors.New("没有处理该操作的权限")
	}

	switch req.Status {
	case 0: // 未操作
		return
	case 1: //同意
		var newMember = models.GroupMemberModel{
			GroupID: groupValidModel.GroupID,
			UserID:  groupValidModel.UserID,
			Role:    0,
		}
		l.svcCtx.DB.Create(&newMember)
	case 2: // 拒绝
	case 3: // 忽略
	case 4: // 删除
		l.svcCtx.DB.Delete(&groupValidModel)
		return
	}
	err = l.svcCtx.DB.Model(&groupValidModel).UpdateColumn("status", req.Status).Error
	if err != nil {
		logx.Errorf("群验证状态更新失败:%s", err)
		return nil, errors.New("群验证状态更新失败")
	}
	return
}
