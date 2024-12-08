package ioc

import "github.com/gin-gonic/gin"

type IocObject interface {
	Name() string
	Init() error
}

// 要单独管理API Handler, API Handler 有个Registry() 注册路由的特色方法
type IocApiObject interface {
	IocObject
	Registry(r gin.IRouter)
}
