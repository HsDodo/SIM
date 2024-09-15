package logic

import (
	"context"
	"errors"
	"fmt"
	"server/chat/api/internal/svc"
	"server/chat/api/internal/types"
	model "server/models/chat"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatDeleteLogic {
	return &ChatDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatDeleteLogic) ChatDelete(req *types.ChatDeleteRequest) (resp *types.NoneResponse, err error) {
	//ChatUserInfo Todo: 删除聊天记录
	// Param: req.ChatIDList []uint, req.UserID uint
	//先查询这个需要删除的聊天记录是否存在，并且是否属于该用户
	var deleteChatIDList []uint
	l.svcCtx.DB.Model(&model.ChatModel{}).Where("id in (?) and (send_user_id = ? or rev_user_id = ?)", req.ChatIDList, req.UserID, req.UserID).Pluck("id", &deleteChatIDList)
	for _, chatID := range req.ChatIDList {
		if NotInList(deleteChatIDList, chatID) {
			logx.Errorf("删除聊天记录失败，被删除的聊天记录中不存在或不属于该用户: %d", chatID)
			return nil, errors.New(fmt.Sprintf("删除聊天记录失败，被删除的聊天记录中不存在或不属于您！"))
		}
	}
	// 再判断这些聊天记录是否已经被删除过
	var deletedIDs []uint
	l.svcCtx.DB.Model(&model.UserChatDeleteModel{}).Where("chat_id in (?) and user_id = ?", deleteChatIDList, req.UserID).Pluck("chat_id", &deletedIDs)
	// 如果已经删除过了，就不再重复删除
	var needDeleteIDs []uint
	if len(deletedIDs) != 0 {
		logx.Errorf("聊天记录已经被删除过了: %v", deletedIDs)
		return nil, errors.New("聊天记录已经被删除过了")
	}
	for _, chatID := range deleteChatIDList {
		if NotInList(deletedIDs, chatID) {
			needDeleteIDs = append(needDeleteIDs, chatID)
		}
	}
	// 上面已经将需要删除的记录过滤完了，接下来就是删除这些记录
	// 删除操作：在user_chat_delete表中插入记录
	var userChatDeleteList []model.UserChatDeleteModel
	for _, chatID := range needDeleteIDs {
		userChatDeleteList = append(userChatDeleteList, model.UserChatDeleteModel{
			UserID: req.UserID,
			ChatID: chatID,
		})
	}

	if len(userChatDeleteList) == 0 {
		logx.Info("没有需要删除的聊天记录")
		return
	}

	err = l.svcCtx.DB.Create(&userChatDeleteList).Error
	if err != nil {
		logx.Errorf("插入删除记录失败: %v", err)
		return nil, err
	}
	return
}

func NotInList(list []uint, id uint) bool {
	for _, v := range list {
		if v == id {
			return false
		}
	}
	return true
}
