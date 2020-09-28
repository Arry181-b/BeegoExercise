package models

/**
* 公用的用于返回发结构体
 */

type Result struct {
	Code int `json:"code"` //接口返回状态类型
	Message string `json:"message"`//接口返回状态对应的描述信息
	Data interface `json:"data"`{} //返回的数据


}
