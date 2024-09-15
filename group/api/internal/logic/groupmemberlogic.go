package logic

import (
	"context"
	"errors"
	"fmt"
	"server/common/list_query"
	"server/common/models"
	group "server/models/group"
	user_rpc "server/user/rpc/userservice"

	"server/group/api/internal/svc"
	"server/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMemberLogic {
	return &GroupMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Data struct {
	GroupID         uint   `gorm:"column:group_id"`
	UserID          uint   `gorm:"column:user_id"`
	Role            int8   `gorm:"column:role"`
	CreatedAt       string `gorm:"column:created_at"`
	MemberNickname  string `gorm:"column:nickname"`
	NewMsgDate      string `gorm:"column:new_msg_date"`
	ProhibitionTime *int   `gorm:"column:forbid_time"`
}

type UserInfo struct {
	UserID   uint32
	Nickname string
	Avatar   string
}

func (l *GroupMemberLogic) GroupMember(req *types.GroupMemberRequest) (resp *types.GroupMemberResponse, err error) {
	// Note: 获取群成员信息列表
	// warning: 可能需要完善，没有验证是否有效
	switch req.Sort {
	case "new_msg_date desc", "new_msg_date asc": // 按照最新发言排序
	case "role asc": //  按照角色升序
	case "created_at desc", "created_at asc": // 按照进群时间排
	default:
		return nil, errors.New("不支持的排序模式")
	}
	column := fmt.Sprintf(fmt.Sprintf("(select group_msg_models.created_at from group_msg_models  where group_msg_models.group_id = %d  and group_msg_models.send_user_id = user_id order by created_at desc limit 1) as new_msg_date", req.GroupID))

	memberList, count, _ := list_query.ListQuery(l.svcCtx.DB, Data{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Sort:  req.Sort,
		Where: l.svcCtx.DB.Where("group_id = ?", req.GroupID),
		Table: func() (string, any) {
			return "(?) as u", l.svcCtx.DB.Model(&group.GroupMemberModel{GroupID: req.GroupID}).
				Select("group_id",
					"user_id",
					"role",
					"created_at",
					"nickname",
					"forbid_time",
					column)
		},
	})

	var userIDList []uint32
	for _, data := range memberList {
		userIDList = append(userIDList, uint32(data.UserID))
	}

	var userInfoMap = map[uint]UserInfo{}
	userListResponse, err := l.svcCtx.UserRpc.UserListInfo(l.ctx, &user_rpc.UserListInfoRequest{
		UserIDs: userIDList,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	for id, info := range userListResponse.UserInfoMap {
		userInfoMap[uint(id)] = UserInfo{
			UserID:   uint32(id),
			Nickname: info.NickName,
			Avatar:   info.Avatar,
		}
	}
	// 查询哪些在线
	onlineUserMap := map[uint32]bool{}
	list, err := l.svcCtx.UserRpc.UserOnlineList(l.ctx, &user_rpc.UserOnlineListRequest{
		UserIds: userIDList,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	for _, id := range list.UserIds {
		onlineUserMap[id] = true
	}
	for _, id := range userIDList {
		if _, ok := onlineUserMap[id]; !ok {
			onlineUserMap[uint32(uint(id))] = false
		}
	}
	// 是否是好友
	friendMap := map[uint32]bool{}
	friendInfoMapRes, err := l.svcCtx.UserRpc.FriendList(l.ctx, &user_rpc.FriendListRequest{
		UserId: uint32(req.UserID),
	})
	for _, info := range friendInfoMapRes.UserInfoMap {
		friendMap[info.UserId] = true
	}
	for _, id := range userIDList {
		if _, ok := friendMap[id]; !ok {
			friendMap[id] = false
		}
	}
	// 封装群员信息
	resp = new(types.GroupMemberResponse)
	for _, member := range memberList {
		resp.List = append(resp.List, types.GroupMemberInfo{
			UserID:          member.UserID,
			UserNickname:    userInfoMap[member.UserID].Nickname,
			Avatar:          userInfoMap[member.UserID].Avatar,
			Role:            member.Role,
			IsOnline:        onlineUserMap[uint32(member.UserID)],
			IsFriend:        friendMap[uint32(member.UserID)],
			ProhibitionTime: member.ProhibitionTime,
		})
	}
	resp.Count = int(count)
	return
}
