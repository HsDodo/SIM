package models

import (
	"server/common/models"
)

type GroupUserRoleModel struct { //用户角色表 0-普通群员 1-管理员 2-群主
	models.Model
	RoleName    string `gorm:"size:32" json:"roleName"`     // 角色名称
	RoleDesc    string `gorm:"size:128" json:"roleDesc"`    // 角色描述
	Permissions string `gorm:"size:256" json:"permissions"` // 权限,用逗号分隔 如 1，2，3，5，6等
}
