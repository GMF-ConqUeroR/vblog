package common

import "github.com/go-playground/validator/v10"

func Validate(obj any) error {
	v := validator.New()
	// 通过为 struct 添加 tag 定义交验规则
	return v.Struct(obj)
}
