package ioc

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var api = map[string]IocApiObject{}

// 对所有的对象统一执行 初始化
// 把所有 api handler 注册root 路由
func InitAllApi(r gin.IRouter) error {
	for k, v := range api {
		err := v.Init()
		if err != nil {
			return fmt.Errorf("init %s error: %s", k, err)
		}
		v.Registry(r)
	}
	return nil
}

func RegisryApi(obj IocApiObject) {
	if _, ok := api[obj.Name()]; ok {
		panic(fmt.Sprintf("%s 已注册!!!", obj.Name()))
	}
	api[obj.Name()] = obj
}

// 获取所有托管对象的名称
func ListApi() (apiNameList []string) {
	for k := range api {
		apiNameList = append(apiNameList, k)
	}
	return apiNameList
}
