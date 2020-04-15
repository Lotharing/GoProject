// Copyright 2020. All rights reserved.
// Author  赵路通
package models

import (
	"GoProject/mappers"
)

//结构体 实现oop  go语言中的"class"
//字母大写则该成员为公有成员(对外可见)，否则是私有成员(对外不可见)
//json序列化操作 'json:"name"'
type User struct {
	UserId     int    `json:"user_id"`
	Account    string `json:"account"`
	Password   string `json:"password"`
	RoleName   string `role_name`
	EmployeeId int    `employee_id`
	UpdateTime string `update_time`
	Status     int    `json:"status"`
}

// 通过用户ID查询用户信息
func FindTeacherInfoById(userId int) (user *User, err error) {
	q := mappers.FetchUserRawSQL(mappers.FindUserById)
	err = Ormer().Raw(q, userId).QueryRow(&user)
	if err != nil {
		return
	}
	return
}

// 通过用户ID更新密码信息
func ResetPwd(userId int, newPwd string) (int64, error) {
	q := mappers.FetchUserRawSQL(mappers.UpdateUserPwd)
	r, err := Ormer().Raw(q, newPwd, userId).Exec()
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}
