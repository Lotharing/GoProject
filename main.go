// Copyright 2020. All rights reserved.
// Author  赵路通
package main

import (
	"GoProject/initial"
	_ "Goproject/routers"
	"github.com/astaxie/beego"
	"net/http"
)

func main() {
	// 初始化数据库
	initial.InitDb()

	// 定义找不到页面
	beego.ErrorHandler("404", func(w http.ResponseWriter, r *http.Request) {
		w.Write(nil)
	})

	beego.Run()
}
