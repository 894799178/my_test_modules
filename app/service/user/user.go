package svr_user

import (
	"errors"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/util/gvalid"
)

const(
	USER_SESSION_MARK = "user_info"
)

var(
	table=g.DB().Table("user").Safe()
)


func SignUp(data g.MapStrStr) error{

	rules :=[]string{
		"passport @required|length:6,16#账号不能为空|账号长度应当在:min到:max之间",
		"password2@required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间",
		"password @required|length:6,16|same:password2#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等",
	}
	if e:=gvalid.CheckMap(data,rules); e!=nil{
		return errors.New(e.String())
	}
	if	_,ok :=data["nickname"]; !ok{
		data["nickname"] = data["passport"]
	}




	return nil
}

//检查张海是否符合规范(目前仅检查唯一性)
func CheckPassport(passport string) bool{
	return QuestUserCountByCondition("passport",passport)
}
//检查昵称是否符合规范
func CheckNickName(nikename string)bool{
	return QuestUserCountByCondition("nickname",nikename)
}

//按照条件进行查找对应的数据是否存在
func QuestUserCountByCondition(conditionName string,value string)bool{
	if i, e := table.Where(conditionName, value).Count(); e!=nil{
		return false
	}else{
		return i == 0
	}
}

