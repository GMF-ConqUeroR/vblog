package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

var (
	config *Config
)

func Values() *Config {
	if config == nil {
		panic("程序配置加载失败，请检查配置文件或者环境变量！！！")
	}
	return config
}

// 加载配置文件
func LoadConfigFromToml(fp string) (*Config, error) {
	conf := DefaultConfig()
	_, err := toml.DecodeFile(fp, conf)
	if err != nil {
		return nil, err
	}
	config = conf
	return conf, nil
}

// 加载环境变量
func LoadConfigFromEnv() (*Config, error) {
	conf := DefaultConfig()
	err := env.Parse(conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
