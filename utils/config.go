package utils

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
)

// 配置文件 yaml
type ConfigEnv struct {
	configFile *yaml.File
}

func getDatabaseConfig() *ConfigEnv {
	return getConfig("database")
}

func getRedisConfig() *ConfigEnv {
	return getConfig("redis")
}

func getConfig(name string) *ConfigEnv {
	filePath := fmt.Sprintf("config/%s.yaml", name)
	return NewEnv(filePath)
}

//配置文件数据解析
func NewEnv(configFile string) *ConfigEnv {
	env := &ConfigEnv{configFile: yaml.ConfigFile(configFile)}
	if env.configFile == nil {
		panic("配置文件打开失败:" + configFile)
	}
	return env
}
