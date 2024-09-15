package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	group "server/models/group"
	models "server/models/user"
	user_rpc "server/user/rpc/userservice"
	"server/utils"
	"time"

	"server/group/api/internal/svc"
	"server/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCreateLogic {
	return &GroupCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupCreateLogic) GroupCreate(req *types.GroupCreateRequest) (resp *types.NoResponse, err error) {
	// todo: add your logic here and delete this line

	var newGroup = group.GroupModel{
		CreatorID:  req.UserID, // 自己创建的群，自己就是群主
		GroupDesc:  "本群创建于" + time.Now().Format("2006-01-02") + ":  群主很懒,什么都没有留下",
		IsSearch:   true,
		VerifyType: 2,
		Size:       50,
	}
	// 获取用户基本信息
	userInfoData, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user_rpc.UserInfoRequest{
		UserId: uint32(req.UserID),
	})
	if err != nil {
		logx.Error(err)
		return nil, errors.New("用户服务错误")
	}

	userInfo := models.UserModel{}
	err = json.Unmarshal(userInfoData.Data, &userInfo)
	if err != nil {
		logx.Error(err)
		return nil, errors.New("用户服务错误")
	}
	if userInfo.UserConf.ForbidCreateGroup {
		return nil, errors.New("当前用户被限制建群")
	}
	groupUserIDList := []uint{req.UserID}
	switch req.Mode {
	case 1: // 直接创建模式
		if req.GroupName == "" {
			return nil, errors.New("群名不能为空")
		}
		if req.Size >= 1000 {
			return nil, errors.New("群规模错误")
		}
		newGroup.GroupName = req.GroupName
		newGroup.Size = req.Size
		newGroup.IsSearch = req.IsSearch
	case 2: // 选人创建模式
		if len(req.UserIDList) == 0 {
			return nil, errors.New("没有要选择的好友")
		}
		for _, id := range req.UserIDList {
			if !utils.InIDsList(groupUserIDList, id) {
				groupUserIDList = append(groupUserIDList, id)
			}
		}
		var friendIDList = []uint32{}
		f, err := l.svcCtx.UserRpc.FriendList(l.ctx, &user_rpc.FriendListRequest{
			UserId: uint32(req.UserID),
		})
		if err != nil {
			logx.Error(err)
			return nil, err
		}
		for friendId, _ := range f.UserInfoMap {
			friendIDList = append(friendIDList, friendId)
		}
		for _, id := range req.UserIDList {
			if !utils.InIDsList(friendIDList, uint32(id)) {
				return nil, errors.New("选择的用户不是好友")
			}
		}

		if req.GroupName == "" {
			newGroup.GroupName = fmt.Sprintf("无名群聊(%d人)", len(groupUserIDList))
		}
	default:
		return nil, errors.New("创建群聊模式错误")
	}
	// 群头像设置, 先默认统一猫猫头
	newGroup.GroupAvatar = "https://img.yzcdn.cn/vant/cat.jpeg"
	err = l.svcCtx.DB.Create(&newGroup).Error
	if err != nil {
		logx.Error(err)
		return nil, errors.New("创建群聊失败")
	}
	// 群成员入库
	var groupMembers []group.GroupMemberModel
	for _, id := range groupUserIDList {
		member := group.GroupMemberModel{
			GroupID: newGroup.ID,
			UserID:  id,
			Role:    0,
		}
		if id == req.UserID {
			member.Role = 2
		}
		groupMembers = append(groupMembers, member)
	}
	l.svcCtx.DB.Create(&groupMembers)
	if err != nil {
		logx.Error(err)
		return nil, errors.New("创建群聊失败")
	}
	return
}
