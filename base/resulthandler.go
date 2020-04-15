// Copyright 2020. All rights reserved.
// 赵路通

package base

import (
	erroresult "GoProject/models/response"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"net/http"
)

const (
	ErrSystemMsg   = "系统错误"
	ErrForbidenMsg = "无访问权限"
)

type ResultHandlerController struct {
	beego.Controller
}

type Result struct {
	Msg    string      `json:"msg"`
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func (c *ResultHandlerController) Success(data interface{}, msg string) {
	c.Data["json"] = Result{Data: data, Status: http.StatusOK, Msg: msg}
	c.ServeJSON()
}

func (c *ResultHandlerController) Failed(msg string) {
	c.Data["json"] = Result{Status: -1, Msg: msg}
	c.ServeJSON()
}

func (c *ResultHandlerController) AbortForbidden() {
	c.Data["json"] = Result{Data: nil, Status: http.StatusForbidden, Msg: ErrForbidenMsg}
	c.ServeJSON()
}

// Handle return http code and body normally, need return
func (c *ResultHandlerController) HandleError(err error) {
	errorResult := &erroresult.ErrorResult{
		Status: -1,
	}
	switch e := err.(type) {
	case *erroresult.ErrorResult:
		errorResult = e
	default:
		if err == orm.ErrNoRows {
			errorResult.Msg = "找不到数据"
		} else {
			logs.Info("返回未知错误：", err.Error())
			errorResult.Msg = ErrSystemMsg
		}
	}

	c.Data["json"] = Result{Data: nil, Status: errorResult.Status, Msg: errorResult.Msg}
	c.ServeJSON()
}
