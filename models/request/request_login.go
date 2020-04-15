// Copyright 2020. All rights reserved.
// Author 赵路通

package request

// 用户名密码登录
type LoginPwdDomain struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
