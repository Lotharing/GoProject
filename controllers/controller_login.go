package controllers

import (
	"GoProject/base"
	"GoProject/models"
)

type LoginController struct {
	base.ResultHandlerController
}

//取用户信息 *方法名首字母大写,要不外部不能调用
func (l *LoginController) GetUserById() {
	//request 中的参数 / default 默认值
	userId, _ := l.GetInt("user_id", 1)
	user, err := models.FindTeacherInfoById(userId)
	if err != nil {
		l.HandleError(err)
		return
	}
	l.Success(user, "获取用户信息成功")
}

//根据用户ID修改密码
func (l *LoginController) UpdPwdById() {
	//request 中的参数 / default 默认值
	userId, _ := l.GetInt("user_id", 1)
	password := l.GetString("password")
	status, err := models.ResetPwd(userId, password)
	if err != nil {
		l.HandleError(err)
		return
	}
	l.Success(status, "成功")
}
