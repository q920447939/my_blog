package controllers

import (
	"my_blog/models"
)

type CategoryController struct {
	BaseResultController
}

func (this *CategoryController) Get ()  {
	this.TplName = "category.html"
	name := this.Input().Get("name")
	if name == "" {
		return
	}
	if err := models.AddCategory(name) ; err != nil{
		this.JsonResult(4000,"保存错误,请稍后重试!")
	}
	return
}

