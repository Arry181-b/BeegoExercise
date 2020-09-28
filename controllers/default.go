package controllers

import "C"
import (
	"BeegoExercise/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}


func (c *MainController) Get() {
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "646844086@qq.com"
	c.TplName = "index.tpl"
	
	//name := c.GetString("name")
	//age1,err := c.GetInt("19")
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

func (c *MainController) Post(){
	//body := c.Ctx.Request.Body
	dataByes, err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收失败，请重试")
		return
	}
	//json包解析
	var person models.Person
	err = json.Unmarshal(dataByes,&person)
	if err != nil{
		fmt.Println(err.Error())
		c.Ctx.WriteString("数据解析失败，try again")
		return
	}
	fmt.Println("用户名：",person.User,"年龄：",person.Age)
	c.Ctx.WriteString("用户名是：" + person.User)
}



//func (c *MainController) Post() {
//	fmt.Println("post类型的请求")
//	user := c.Ctx.Request.FormValue("user")
//	fmt.Println("用户名为：", user)
//	psd := c.Ctx.Request.FormValue("psd")
//	fmt.Println("密码是",psd)

	////与固定值比较
	//if user != "admin" || psd != "20010331" {
	//	失败页面
	//	c.Ctx.ResponseWriter.Write([]byte("对不起，数据不正确"))
	//	return
	//}
	//
	//c.Ctx.ResponseWriter.Write([]byte("恭喜，数据正确"))
	//}











