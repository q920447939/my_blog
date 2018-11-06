package models

import (
	"errors"
	"my_blog/connection/mysql"
	"time"
)

// 分类
type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

func FindCategoryByName(name string) (e error)  {
	category := make([]Category,10)
	if err := mysql.Db.Unscoped().Table("Category").Where("title = ?", name).Find(category).Error ;err != nil {
		e = errors.New("查询数据失败")
	}
	return
}

func AddCategory(name string ) error  {
	category := Category{
		Title:           name,
		Created:         time.Now(),
		TopicTime:       time.Now(),
		TopicLastUserId: 1,
	}
	if err := mysql.Db.Create(&category).Error ; err != nil {
		return err
	}
	return  nil
}