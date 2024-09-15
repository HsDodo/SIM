package list_query

import (
	"fmt"
	"gorm.io/gorm"
	"server/common/models"
)

type Option struct {
	PageInfo models.PageInfo
	Where    *gorm.DB
	Debug    bool
	Joins    string
	Likes    map[string]string    // 模糊匹配的字段
	Preload  []string             // 预加载字段
	Table    func() (string, any) // 子查询
	Groups   []string             // 分组
	Sort     string               // 排序
}

func ListQuery[T any](db *gorm.DB, model T, option Option) (list []T, count int64, err error) {
	if option.Debug {
		db = db.Debug()
	}
	query := db.Where(model)
	// 模糊匹配
	if len(option.Likes) > 0 {
		likeQuery := db.Where("1 = 0")
		for column, value := range option.Likes {
			likeQuery.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", value))
		}
		query.Where(likeQuery)
	}

	if option.Table != nil {
		table, data := option.Table()
		query = query.Table(table, data)
	}

	if option.Joins != "" {
		query = query.Joins(option.Joins)
	}

	if option.Where != nil {
		query = query.Where(option.Where)
	}
	if len(option.Groups) > 0 {
		for _, group := range option.Groups {
			query = query.Group(group)
		}
	}

	// 求总数
	query.Model(model).Count(&count)
	// 预加载
	for _, preload := range option.Preload {
		query = query.Preload(preload)
	}
	// 分页查询
	if option.PageInfo.Page <= 0 {
		option.PageInfo.Page = 1
	}
	if option.PageInfo.Limit != -1 { // -1 代表不分页，查全部
		if option.PageInfo.Limit <= 0 {
			option.PageInfo.Limit = 10
		}
	}
	offset := (option.PageInfo.Page - 1) * option.PageInfo.Limit
	if option.Sort != "" {
		query = query.Order(option.Sort)
	}

	err = query.Limit(option.PageInfo.Limit).Offset(offset).Find(&list).Error
	return
}
