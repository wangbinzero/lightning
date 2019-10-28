package initializers

import (
	"github.com/labstack/echo"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {

	}
}

//语言选择
func chooseLanguage(content echo.Context) {
	var language string
	var lqs []local.L
}
