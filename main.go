package main

import (
	"flag"
	"gorm.io/gorm"
	"log"
	"server/core"
	logs "server/logs/model"
	chat "server/models/chat"
	file "server/models/file"
	group "server/models/group"
	msg "server/models/message"
	user "server/models/user"
)

type Options struct {
	// The name of the package to generate the code in.
	DB bool
}

func main() {
	var opt Options
	var db *gorm.DB
	flag.BoolVar(&opt.DB, "db", false, "Generate database code") //value of db is false, if -db is not set
	flag.Parse()
	if opt.DB {
		// 初始化Mysql数据库
		// InitGrom()
		db = core.InitGorm("root:root@tcp(127.0.0.1:3306)/sim_db?charset=utf8&parseTime=True&loc=Local")
		err := db.AutoMigrate(
			&user.UserModel{},
			&user.UserConfModel{},
			&msg.MsgModel{},
			&group.GroupModel{},
			&group.GroupMemberModel{},
			&group.GroupMsgModel{},
			&group.GroupVerifyModel{},
			&user.UserRoleModel{},
			&group.GroupUserRoleModel{},
			&user.FriendshipModel{},
			&file.FileModel{},
			&user.FriendVerifyModel{},
			// 聊天模型
			&chat.ChatModel{},           // 聊天记录
			&chat.UserChatDeleteModel{}, // 用户删除聊天记录
			&chat.TopUserModel{},        // 置顶用户
			// 日志模型
			&logs.LogModel{},
		)
		if err != nil {
			log.Fatalln("自动创建表失败: ", err)
		}
		log.Println("自动创建表成功!")
	}

}
