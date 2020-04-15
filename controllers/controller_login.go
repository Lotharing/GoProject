package controllers

import (
	"GoProject/base"
	"GoProject/models"
)

type LoginController struct {
	base.ResultHandlerController
}

//取用户信息
func (l *LoginController) Logout() {
	userId, _ := l.Ctx.Input.GetData("user_id").(int)
	user, err := models.FindTeacherInfoById(userId)
	if err != nil {
		l.HandleError(err)
		return
	}
	l.Success(user, "获取用户信息成功")
}
