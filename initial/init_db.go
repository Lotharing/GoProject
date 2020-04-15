// Copyright 2020. All rights reserved.
// Author  赵路通

package initial

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	DbAlias      = "default"
	DbDriverName = "mysql"
	maxIdle      = 10
	maxOpen      = 20
)

func InitDb() {
	orm.RegisterDataBase(DbAlias, DbDriverName, beego.AppConfig.String("db.datasource"), maxIdle, maxOpen)

	db, err := orm.GetDB()
	if err != nil {
		panic(err)
	}
	ttl := beego.AppConfig.DefaultInt("db.connTTL", 30)

	db.SetConnMaxLifetime(time.Duration(ttl) * time.Second)

	orm.Debug = beego.AppConfig.DefaultBool("db.showSql", false)
}
