package controllers

import "C"
import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//获取get类型请求方式
	name := c.Ctx.Input.Query("name")
	age := c.Ctx.Input.Query("age")
	sex := c.Ctx.Input.Query("sex")
	fmt.Println(name,age,sex)

	//以admin,19为正确数据验证
	if name != "admin" || age !=  "18" || sex != "female"{
		c.Ctx.ResponseWriter.Write([]byte("数据验证错误"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据验证成功"))

}



func (c *MainController) Post() {
	fmt.Println("post类型的请求")
	user := c.Ctx.Request.FormValue("user")
	fmt.Println("用户名为：", user)
	psd := c.Ctx.Request.FormValue("psd")
	fmt.Println("密码是",psd)

	//与固定值比较
	if user != "admin" || psd != "20010331" {
		//数百页面
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据不正确"))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("恭喜，数据正确"))

	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "646844086@qq.com"
	c.TplName = "index.tpl"
}











