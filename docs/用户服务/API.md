### 用户服务 API

#### 主要功能

> 1. 获取用户信息 `get` /api/user/user_info 
> 2. 修改用户信息 `put` /api/user/user_info
> 3. 获取好友信息 `get` /api/user/friend_info
> 4. 获取好友列表 `get` /api/user/friend_list
> 5. 修改好友备注 `put` /api/user/friend
> 6. 搜索用户      `get` /api/user/search
> 7. 获取好友验证信息 `get` /api/user/valid
> 8. 添加好友      `post` /api/user/friend
> 9. 验证好友请求 `put`  /api/user/valid_status
> 10. 删除好友      `delete` /api/user/friends

#### 用户服务 API

```go
syntax = "v1"

info (
	title:   "用户服务"
	desc:    "用户服务api"
	author:  "sen"
	version: "v0.01"
)

type GetUserInfoRequest {
	UserID uint `header:"userID"`
}

type UserInfoResponse {
	data []byte `json:"data"`
}

//用户信息修改
type UpdateUserInfoRequest {
	//基本用户信息
	UserID   uint    `header:"userID" `
	NickName *string `json:"nickname,optional" user:"nickname"`
	Avatar   *string `json:"avatar,optional" user:"avatar"`
	Abstract *string `json:"abstract,optional" user:"abstract"`
	// 配置信息
	RecallMessage        *string `json:"recallMessage,optional" userConf:"recallMessage"`
	FriendOnline         *bool   `json:"friendOnline,optional" userConf:"friendOnline"`
	Sound                *bool   `json:"sound,optional" userConf:"sound"`
	SecureLink           *bool   `json:"secureLink,optional" userConf:"secureLink"`
	SavePwd              *bool   `json:"savePwd,optional" userConf:"savePwd"`
	SearchUser           *int8   `json:"searchUser,optional" userConf:"searchUser"`
	Verification         *int8   `json:"verification,optional" userConf:"verification"`
	VerificationQuestion *string `json:"verificationQuestion,optional" userConf:"verificationQuestion"`
}

type UpdateUserInfoResponse {}

type FriendInfoRequest {
	UserID   uint `header:"userID"`
	FriendID uint `form:"friendID"`
}

type FriendInfoResponse {
	NickName string `json:"nickname"` // 好友昵称
	Avatar   string `json:"avatar"` // 好友头像
	Abstract string `json:"abstract"` // 好友简介
	Alias    string `json:"alias"` // 好友备注
}

type GetFriendListRequest {
	UserID   uint `header:"userID"`
	Page     int  `form:"page"`
	PageSize int  `form:"pageSize"`
}

type GetFriendListResponse {
	List  []FriendInfoResponse `json:"list"`
	Count uint                 `json:"count"`
}

type UpdateFriendAliasRequest {
	UserID   uint   `header:"userID"`
	FriendID uint   `json:"friendID"`
	Alias    string `json:"alias"`
}

type SearchUserRequest {
	UserID   uint   `header:"userID"`
	Key      string `form:"key, optional"` // 用户id和昵称搜索
	Online   bool   `form:"online, optional"` // 是否在线
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
}

type SearchInfo {
	UserID   uint   `json:"userID"`
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Abstract string `json:"abstract"`
	IsFriend bool   `json:"isFriend"`
}

type SearchUserResponse {
	List  []SearchInfo `json:"list"`
	Count uint         `json:"count"`
}

type UpdateFriendAliasResponse {}

type UserValidRequest {
	UserID   uint `header:"userID"`
	FriendID uint `form:"friendID"`
}

type UserValidResponse {
	Verification         int8                 `json:"verification"`
	VerificationQuestion VerificationQuestion `json:"verificationQuestion"`
}

type VerificationQuestion {
	Problem1 *string `json:"problem1,optional" user_conf:"problem1"`
	Problem2 *string `json:"problem2,optional" user_conf:"problem2"`
	Problem3 *string `json:"problem3,optional" user_conf:"problem3"`
	Answer1  *string `json:"answer1,optional" user_conf:"answer1"`
	Answer2  *string `json:"answer2,optional" user_conf:"answer2"`
	Answer3  *string `json:"answer3,optional" user_conf:"answer3"`
}

type AddFriendRequest {
	UserID               uint                  `header:"userID"`
	FriendID             uint                  `json:"friendID"`
	VerifyMsg            string                `json:"verifyMsg, optional"` // 验证消息
	VerificationQuestion *VerificationQuestion `json:"verificationQuestion,optional"` // 问题和答案
}

type FriendValidStatusRequest {
	UserID  uint `header:"userID"`
	VaildID uint `json:"vaildID"`
	Status  int8 `json:"status"`
}

type FriendDeleteRequest {
	UserID   uint `header:"userID"`
	FriendID uint `json:"friendID"`
}

type FriendDeleteResponse {}

type FriendValidStatusResponse {}

type AddFriendResponse {}

service user {
	@handler getUserInfo
	get /api/user/user_info (GetUserInfoRequest) returns (UserInfoResponse) //获取用户信息

	@handler updateUserInfo
	put /api/user/user_info (UpdateUserInfoRequest) returns (UpdateUserInfoResponse) //修改用户信息

	@handler getFriendInfo
	get /api/user/friend_info (FriendInfoRequest) returns (FriendInfoResponse) //获取好友信息

	@handler getFriendList
	get /api/user/friend_list (GetFriendListRequest) returns (GetFriendListResponse) //获取好友列表

	@handler updateFriendAlias
	put /api/user/friend (UpdateFriendAliasRequest) returns (UpdateFriendAliasResponse) //修改好友备注

	@handler searchUser
	get /api/user/search (SearchUserRequest) returns (SearchUserResponse) //搜索用户

	@handler userValid
	post /api/user/valid (UserValidRequest) returns (UserValidResponse) //获取好友验证信息

	@handler addFriend
	post /api/user/friend (AddFriendRequest) returns (AddFriendResponse) //添加好友

	@handler validStatus
	put /api/user/valid_status (FriendValidStatusRequest) returns (FriendValidStatusResponse) //验证好友请求

	@handler friendDelete
	delete /api/user/friends (FriendDeleteRequest) returns (FriendDeleteResponse) // 删除好友
}

// goctl api go --api user.api --dir .

```

