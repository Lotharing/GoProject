// Copyright 2020. All rights reserved.
// Author  赵路通
package main

import (
	"GoProject/initial"
	"github.com/astaxie/beego"
)

func main() {
	// 初始化数据库
	initial.InitDb()

	beego.Run()
}
