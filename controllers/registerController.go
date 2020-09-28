package controllers

import (
	"BeegoProject/db_mysql"
	"BeegoProject/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post() {
	dataBytes, err := ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil {
		return
	}
}

var user models.User
err = json.Unmarshal(dataBytes, &user)
if err != nil {
	r.Ctx.WriteString("数据解析错误请重试")
}

row, err := db_mysql.Adduser(user)
if err != nil {
	r.CtxWriteString("注册信息失败，请重试")
	return

}
fmt.Println(row)
r.Ctx.WriteString("恭喜，注册用户信息成功")

//
//if err ! nil {
//	//r.Ctx.WriteString("数据解析错误，请重试")
//	result：= models.Result {
//		Code:0,
//		Message:"数据解析错误，请重试"，
//		Data:nil,
//}
//
//	r.Data["json"] = &result
//	r.ServeJson()
//	return
//}

