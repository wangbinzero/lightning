package main

import (
	newrelic "github.com/dafiti/echo-middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	envConfig "lightning/config"
	"lightning/routers"
)

func main() {
	e := echo.New()

	if envConfig.CurrentEnv.Newrelic.AppName != "" && envConfig.CurrentEnv.Newrelic.LicenseKey != "" {
		e.Use(newrelic.NewRelic(envConfig.CurrentEnv.Newrelic.AppName, envConfig.CurrentEnv.Newrelic.LicenseKey))
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	routers.SetV1Interfaces(e)
	e.Start(":12345")
}

//初始化配置信息
func initialize() {
	envConfig.InitEnv()
}
