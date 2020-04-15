// Copyright 2020. All rights reserved.
// Author 赵路通
package models

import (
	"github.com/astaxie/beego/orm"
	"sync"
)

var (
	globalOrm orm.Ormer
	once      sync.Once
)

func Ormer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}
