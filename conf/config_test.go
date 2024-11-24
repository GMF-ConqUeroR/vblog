package conf_test

import (
	"testing"

	"gitee.com/xpigpig/vblog/conf"
	"github.com/BurntSushi/toml"
)

// /Volumes/ConqueroR/Golang/vBlog/conf/config_test.go:16: &{0xc0000761e0}
// 由于以上测试结果可读性差（&{0xc0000761e0}），可以通过实现一个Stringger接口的方法, fmt包使用,
// 凡是使用了fmt进行打印的，都会按照自定义格式打印
func TestDecodeConfigFromTomlFile(t *testing.T) {
	confObj := &conf.Config{}
	_, err := toml.DecodeFile("test/config.toml", confObj)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(confObj)
}

func TestLoadConfigFromToml(t *testing.T) {
	// 程序启动时先加载配置，然后通过全局变量访问
	_, err := conf.LoadConfigFromToml("test/config.toml")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(conf.Values())
}
