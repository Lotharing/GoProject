// 分页工具类
package utils

import (
	"regexp"
	"strings"
)

type Page struct {
	Page      int         `json:"page"`
	Size      int         `json:"size"`
	Total     int         `json:"total"`
	TotalPage int         `json:"total_page"`
	Items     interface{} `json:"items"`
}

func PageInfo(count int, pageNo int, pageSize int, list interface{}) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Page{Page: pageNo, Size: pageSize, TotalPage: tp, Total: count, Items: list}
}

// 返回查询总数SQL
func ForamtCountSubSQL(text string) string {
	reg := regexp.MustCompile(`(?iU:SELECT.*FROM)`)
	r := reg.ReplaceAllString(strings.Replace(text, "\n", " ", -1), "SELECT COUNT(1) FROM")
	return r
}

// 返回查询总数SQL
func ForamtCountSQL(text string) string {
	reg := regexp.MustCompile(`(?iU:SELECT).*FROM`)
	r := reg.ReplaceAllString(strings.Replace(text, "\n", " ", -1), "SELECT COUNT(1) FROM")
	return r
}

// 拼凑查询总数SQL语句参数
func GetFormatCountSQL(text string, filter string) string {
	countSql := ForamtCountSQL(text) + filter
	return countSql
}
