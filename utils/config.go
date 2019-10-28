package utils

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"log"
	"strconv"
	"time"
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

//读取数据类型
func (env *ConfigEnv) Get(spec, defaultValue string) string {
	value, err := env.configFile.Get(spec)
	if err != nil {
		value = defaultValue
	}
	return value
}

func (env *ConfigEnv) GetInt(spec string, defaultValue int) int {
	str := env.Get(spec, "")
	if str == "" {
		return defaultValue
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		log.Panic("数据类型转换错误 int:", spec, str)
	}
	return val
}

func (env *ConfigEnv) GetDuration(spec string, defaultValue string) time.Duration {
	str := env.Get(spec, "")
	if str == "" {
		str = defaultValue
	}
	duration, err := time.ParseDuration(str)
	if err != nil {
		log.Panic("数据类型转换错误 duration:", spec, str)
	}
	return duration
}
