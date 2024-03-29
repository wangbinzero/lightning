package initializers

import (
	"github.com/labstack/echo"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"lightning/utils"
	"log"
	"path/filepath"
)

type ApiInterface struct {
	Method                string `yaml:"method"`
	Path                  string `yaml:"path"`
	Auth                  bool   `yaml:"auth"`
	Sign                  bool   `yaml:"sign"`
	CheckTimestamp        bool   `yaml:"check_timestamp"`
	LimitTrafficWithIp    bool   `yaml:"limit_traffic_with_ip"`
	LimitTrafficWithEmail bool   `yaml:"limit_traffic_with_email"`
	IsRabbitMqConnected   bool   `yaml:"is_rabbitmq_connected"`
}

var GlobalApiInterfaces []ApiInterface

func LoadInterfaces() {
	path_str, _ := filepath.Abs("config/interfaces.yaml")
	content, err := ioutil.ReadFile(path_str)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = yaml.Unmarshal(content, &GlobalApiInterfaces)
	if err != nil {
		log.Fatal(err)
	}
}

//api授权
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		params := make(map[string]string)
		for k, v := range context.QueryParams() {
			params[k] = v[0]
		}
		values, _ := context.FormParams()
		for k, v := range values {
			params[k] = v[0]
		}
		context.Set("params", params)
		var currentApiInterface ApiInterface
		for _, apiInterface := range GlobalApiInterfaces {
			if context.Path() == apiInterface.Path && context.Request().Method == apiInterface.Method {
				currentApiInterface = apiInterface
				if currentApiInterface.LimitTrafficWithEmail && LimitTrafficWithEmail(context) != true {
					return utils.BuildError("3043")
				}

				//TODO limit_with_ip

				if apiInterface.Auth != true {
					return next(context)
				}
			}
		}

		if currentApiInterface.Path==""{
			return utils.BuildError("3025")
		}
		if context.Request().Header.Get("Authorization")==""{
			return utils.BuildError("4001")
		}
		if currentApiInterface.CheckTimestamp&&checkTimestamp(context,&params)==false{
			return utils.BuildError("3050")
		}
		//TODO 校验 MQ 连接
		//if currentApiInterface.IsRabbitMqConnected&&!Isra

		db:=utils.MainDbBegin()
		defer db.DbRollback()
		var
		return nil

	}
}

//TODO 语言选择
//func chooseLanguage(content echo.Context) {
//	var language string
//	var lqs []local
//}
