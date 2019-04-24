package svr_user

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/gogf/gf/g/util/gvalid"
)

const (
	USER_SESSION_MARK = "user_info"
)

var (
	table = g.DB().Table("user").Safe()
)

//用户注册
func SignUp(data g.MapStrStr) error {
	//数据校验
	rules := []string{
		"passport @required|length:6,16#账号不能为空|账号长度应当在:min到:max之间",
		"password2@required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间",
		"password @required|length:6,16|same:password2#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等",
	}
	if e := gvalid.CheckMap(data, rules); e != nil {
		return errors.New(e.String())
	}
	if _, ok := data["nickname"]; !ok {
		data["nickname"] = data["passport"]
	}
	//唯一性数据检查
	if !CheckPassport(data["passport"]) {
		return errors.New(fmt.Sprintf("帐号 %s 已存在", data["passport"]))
	}

	if !CheckNickName(data["nickname"]) {
		return errors.New(fmt.Sprintf("昵称 %s 已存在", data["nickname"]))
	}
	//记录帐号创建时间/注册时间
	if _, ok := data["create_time"]; !ok {
		data["create_time"] = gtime.Now().String()
	}

	if _, err := table.Filter().Data(data).Save(); err != nil {
		return err
	}
	return nil
}

//判断是否已经登录
func IsSignedIn(session *ghttp.Session) bool {
	return session.Contains(USER_SESSION_MARK)
}

//用户登出
func SignOut(s *ghttp.Session) {
	s.Remove(USER_SESSION_MARK)
}

//用户登录
func SignIn(passport, password string, session *ghttp.Session) error {
	record, err := table.Where("passport =? and password =?", passport, password).One()
	if err != nil {
		return err
	}
	if record == nil {
		return errors.New("帐号或密码错误")
	}
	session.Set(USER_SESSION_MARK, record)
	return nil

}

//检查张海是否符合规范(目前仅检查唯一性)
func CheckPassport(passport string) bool {
	return QuestUserCountByCondition("passport", passport)
}

//检查昵称是否符合规范
func CheckNickName(nikename string) bool {
	return QuestUserCountByCondition("nickname", nikename)
}

//按照条件进行查找对应的数据是否不存在
func QuestUserCountByCondition(conditionName string, value string) bool {
	if i, e := table.Where(conditionName, value).Count(); e != nil {
		return false
	} else {
		return i == 0
	}
}
