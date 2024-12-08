package ioc

import "fmt"

// 管理Controller对象
// 采用一个map来存储对象, 存储在内存中, 程序启动 解决多个impl对象的依赖问题
var controller = map[string]IocObject{}

// 对注册的所有的对象统一执行初始化
func InitAllControllers() error {
	for k, v := range controller {
		err := v.Init()
		if err != nil {
			return fmt.Errorf("init %s error: %s", k, err)
		}
	}
	return nil
}

// 注册需要托管的对象
func RegistryController(obj IocObject) {
	if _, ok := controller[obj.Name()]; ok {
		panic(fmt.Sprintf("%s 已注册!!!", obj.Name()))
	}
	controller[obj.Name()] = obj
}

// 获取指定的托管对象
func GetController(name string) any {
	v, ok := controller[name]
	if !ok {
		panic(fmt.Sprintf("%s 不存在!!!", name))
	}
	return v
}

// 获取所有托管对象的名称
func ListContrillers() (controllerNameList []string) {
	for k := range controller {
		controllerNameList = append(controllerNameList, k)
	}
	return controllerNameList
}
