package main

import (
	_ "BeegoExercise/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//定义config变量，接收并赋值为全局配置变量
	config := beego.AppConfig
	//或者配置选项
	appname := config.String("appname")
	fmt.Println("应用项目名称：",appname)
	config.Int("httpport")
	port,err := config.Int("httpport")
	if err != nil {
		//配置信息解析错误
		panic("项目配置信息解析错误，请查验后重试")
	}
	fmt.Println("应用监听端口：",port)
	//driver := config.String("db_driver") //数据库驱动
	//dbUser := config.String("db_user")//数据库用户名
	//dbPassword := config.String("db_password") //密码
	//dbIP := config.String("db_ip")
	//dbName := config.String("db_name")

	//db,err := sql.Open(driver,dbUser+":"+dbPassword+"@tcp("+dbIP+")/"+dbName+"?charset=utf8")
	//if err != nil {
	//	panic("数据库连接失败，请重试")
	//}
	//fmt.Println(db)

	beego.Run()
}

