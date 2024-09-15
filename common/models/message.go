package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type MsgType int8

const (
	TextMsgType MsgType = iota + 1
	ImageMsgType
	VideoMsgType
	FileMsgType
	VoiceMsgType
	VoiceCallMsgType
	VideoCallMsgType
	WithdrawMsgType
	ReplyMsgType
	QuoteMsgType
	AtMsgType
	TipMsgType
	FriendOnlineMsgType
	ImageTextMsgType
)

type SystemMsg struct {
	Type    int8    `json:"type"`    // 消息类型 0-正常消息 1-涉黄 2-涉政 3-广告 4-违法 5-其他
	Content *string `json:"content"` // 消息内容
}

func (s *SystemMsg) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), s) //将 msg的数据解析到data中
}
func (s SystemMsg) Value() (driver.Value, error) {
	marshal, err := json.Marshal(s)
	return string(marshal), err //将msg的数据转为json格式，再以string格式存入数据库
}

type Message struct {
	Type             MsgType              `json:"type"` // 消息类型 1-文本 2-图片 3-文件 4-音频 5-视频 6-语音通话 7-视频通话 8-撤回消息 9-转发消息 10-回复消息 11-@消息 12-提示消息 13-好友上线提醒 14-图文消息
	TextMessage      *string              `json:"textMsg"`
	ImageMessage     *ImageMessage        `json:"imageMsg"`        // 片消息
	FileMessage      *FileMessage         `json:"fileMsg"`         // 文件消息
	AudioMessage     *AudioMessage        `json:"audioMsg"`        // 音频消息
	VideoMessage     *VideoMessage        `json:"videoMsg"`        // 视频消息
	VoiceCallMessage *VoiceCallMessage    `json:"voiceMsg"`        // 语音通话消息
	VideoCallMessage *VideoCallMessage    `json:"videoCallMsg"`    // 视频通话消息
	WithdrawMessage  *WithdrawMessage     `json:"withdrawMsg"`     // 撤回消息
	ForwardMessage   *ForwardMessage      `json:"forwardMsg"`      // 转发消息
	ReplyMessage     *ReplyMessage        `json:"replyMsg"`        // 回复消息
	AtMessage        *AtMessage           `json:"atMsg"`           // @消息
	TipMessage       *TipMessage          `json:"tipMsg"`          // 提示消息
	FriendOnlineMsg  *FriendOnlineMessage `json:"friendOnlineMsg"` // 好友上线提醒
	ImageTextMessage *ImageTextMessage    `json:"imageTextMsg"`    // 图文消息
}

func (msg Message) MsgPreview() string {
	switch msg.Type {
	case 1:
		return *msg.TextMessage
	case 2:
		return "[图片消息] - " + msg.ImageMessage.Title
	case 3:
		return "[文件消息] - " + msg.FileMessage.Title
	case 4:
		return "[音频消息]"
	case 5:
		return "[视频消息] - " + msg.VideoMessage.Title
	case 6:
		return "[语音通话]"
	case 7:
		return "[视频通话]"
	case 8:
		return "[撤回消息] - " + *msg.WithdrawMessage.OriginMsg.TextMessage
	case 9:
		return "[转发消息]"
	case 10:
		return "[回复消息] - " + msg.ReplyMessage.Content
	case 11:
		return "[@消息] - " + msg.AtMessage.Content
	case 12:
	case 13:
	case 14:
		return "[图文消息]-" + msg.ImageTextMessage.Content

	}
	return *msg.TextMessage
}

func (msg *Message) Scan(data interface{}) error { //文件信息可能较大，用指针接收者
	return json.Unmarshal(data.([]byte), msg) //将 msg的数据解析到data中
}
func (msg Message) Value() (driver.Value, error) {
	marshal, err := json.Marshal(msg)
	return string(marshal), err //将msg的数据转为json格式，再以string格式存入数据库
}

type ImageMessage struct {
	Title string `gorm:"size:128" json:"title"` // 图片标题
	Src   string `gorm:"size:128" json:"src"`   // 图片地址
}
type FileMessage struct {
	Title string `gorm:"size:128" json:"title"` // 文件标题
	Src   string `gorm:"size:128" json:"src"`   // 文件地址
	Size  int    `json:"size"`                  // 文件大小
	Type  string `gorm:"size:32" json:"type"`   // 文件类型
}
type AudioMessage struct {
	Src  string `gorm:"size:128" json:"src"` // 音频地址
	Time int    `json:"time"`                // 音频时长
}
type VideoMessage struct {
	Title string `gorm:"size:128" json:"title"` // 视频标题
	Src   string `gorm:"size:128" json:"src"`   // 视频地址
	Time  int    `json:"time"`                  // 视频时长
}
type VoiceCallMessage struct {
	StartTime time.Time `json:"startTime"` // 通话开始时间
	EndTime   time.Time `json:"endTime"`   // 通话结束时间
	EndReason int8      `json:"endReason"` // 通话结束原因 0-正常结束 1-对方未接听 2-对方拒绝接听 3-对方忙线 4-对方不在线 5-对方已挂断 6-自己已挂断 7-网络异常
}
type VideoCallMessage struct {
	StartTime  time.Time `json:"startTime"`  // 通话开始时间
	EndTime    time.Time `json:"endTime"`    // 通话结束时间
	EndReason  int8      `json:"endReason"`  // 通话结束原因 0-正常结束 1-对方未接听 2-对方拒绝接听 3-对方忙线 4-对方不在线 5-对方已挂断 6-自己已挂断 7-网络异常
	ClientFlag int8      `json:"clientFlag"` // 标识，标识客户端弹框的模式
	ServerFlag int8      `json:"serverFlag"` // 服务器处理逻辑
	Msg        string    `json:"msg"`        // 提示消息
	Type       string    `json:"type"`       // WebRTC 消息类型 create_offer create_answer create_candidate 等
	Data       any       `json:"data"`       // 附加数据如 offer answer candidate
}
type WithdrawMessage struct {
	// 撤回消息
	MsgID     uint     `json:"msgId"` // 撤回的消息ID
	OriginMsg *Message `json:"-"`     // 原消息
}
type ForwardMessage struct { // 转发消息
	MsgID uint `json:"msgId"` // 转发的消息ID
}
type ReplyMessage struct {
	MsgID             uint      `json:"msgId"`           // 回复的消息ID
	Content           string    `json:"content"`         // 回复的消息内容, 目前只限回复文本消息
	ReplyUserID       uint      `json:"userID"`          // 回复的用户ID
	ReplyUserNickName string    `json:"nickName"`        // 回复的用户昵称
	OriginMsgDate     time.Time `json:"originMsgDate"`   // 原消息时间
	Msg               *Message  `json:"replyMsg"`        // 回复的消息
	ReplyMsgPreview   string    `json:"replyMsgPreview"` // 回复的消息预览
}
type AtMessage struct {
	UserID  uint     `json:"userID"`    // @的用户ID
	Content string   `json:"content"`   // @的内容
	Msg     *Message `json:"atMessage"` // @的消息
}

type TipMessage struct {
	Status  string `json:"status"`  // 提示状态 error success warning info
	Content string `json:"content"` // 提示内容
}

type FriendOnlineMessage struct {
	NickName string `json:"nickName"` // 好友昵称
	Avatar   string `json:"avatar"`   // 好友头像
	Content  string `json:"content"`  // 在线状态
	FriendID uint   `json:"friendID"` // 好友ID
}

type ImageTextMessage struct {
	Content string `json:"content"` // 带有html标签的文本
}

func (t ImageMessage) Validate() error {
	if t.Title == "" {
		return errors.New("请输入标题")
	}
	if t.Src == "" {
		return errors.New("请输入图片地址")
	}
	return nil
}

func (video VideoMessage) Validate() error {
	if video.Title == "" {
		return errors.New("请输入标题")
	}
	if video.Src == "" {
		return errors.New("请输入视频地址")
	}
	return nil
}

func (audio AudioMessage) Validate() error {
	if audio.Src == "" {
		return errors.New("请输入音频地址")
	}
	return nil
}

func (file FileMessage) Validate() error {

	if file.Src == "" {
		return errors.New("请输入文件地址")
	}
	return nil
}

// 判断消息格式是否正确
func (msg Message) IsValid() error {
	switch msg.Type {
	case TextMsgType:
		if msg.TextMessage == nil {
			return errors.New("文本消息不能为空")
		}
		if *(msg.TextMessage) == "" {
			return errors.New("请输入文本消息")
		}
	case ImageMsgType:
		if msg.ImageMessage == nil {
			return errors.New("图片消息不能为空")
		}
		return msg.ImageMessage.Validate()
	case VideoMsgType:
		if msg.VideoMessage == nil {
			return errors.New("视频消息不能为空")
		}
		return msg.VideoMessage.Validate()
	case FileMsgType:
		if msg.FileMessage == nil {
			return errors.New("文件消息不能为空")
		}
		return msg.FileMessage.Validate()
	}
	return nil
}
