package models

import (
	"gorm.io/gorm"
)

type FriendshipModel struct {
	gorm.Model
	UserID   uint   `gorm:"not null" json:"user_id"`
	FriendID uint   `gorm:"not null" json:"friend_id"`
	Accepted bool   `gorm:"default:false" json:"accepted"`
	Alias    string `json:"alias"`
}

func (friendship *FriendshipModel) IsFriend(db *gorm.DB, userID, friendID uint) bool {
	// 判断两个人是否是好友
	err := db.Model(&FriendshipModel{}).
		Where("(user_id = ? AND friend_id = ?) or (user_id = ? and friend_id = ?) AND accepted = ?", userID, friendID, friendID, userID, true).
		First(friendship).Error
	if err != nil {
		return false
	}
	return true
}

func GetFriendshipIDs(db *gorm.DB, userID uint) (friendIDs []uint, err error) {
	// 获取用户的所有好友ID
	err = db.Model(&FriendshipModel{}).Where("user_id = ? AND accepted = ?", userID, true).Pluck("friend_id", &friendIDs).Error
	if err != nil {
		return
	}

	additionalFriendIDs := []uint{}
	err = db.Model(&FriendshipModel{}).Where("friend_id = ? AND accepted = ?", userID, true).Pluck("user_id", &additionalFriendIDs).Error
	if err != nil {
		return
	}
	friendIDs = append(friendIDs, additionalFriendIDs...)
	return
}
