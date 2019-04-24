/*
@Time : 2019/4/24 14:33
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package router

import (
	"github.com/gogf/gf/g"
	"my_test_modules/app/api/user"
)

func init() {
	g.Server().BindObject("/user", new(api_user.Controller))
}
