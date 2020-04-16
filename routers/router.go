package routers

import (
	"GoProject/controllers"
	"github.com/astaxie/beego"
)

var (
	//用户信息操作控制器
	loginController = new(controllers.LoginController)
)

func init() {

	//可以使用拦截beego.InsertFilter()

	ns := beego.NewNamespace("/api",
		// 登录相关路由配置
		beego.NSNamespace("/v1/login",
			beego.NSRouter("/getUserById", loginController, "post:GetUserById"),
			beego.NSRouter("/updPwdById", loginController, "post:UpdPwdById"),
		),
	)
	beego.AddNamespace(ns)
}
