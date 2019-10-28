package main

import (
	"fmt"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			cc := &CustomContext{context}
			return handlerFunc(cc)
		}
	})
	e.GET("/", func(context echo.Context) error {
		cc := context.(*CustomContext)
		cc.Foo()
		cc.Bar()
		return cc.String(200, "ok")
	})

	e.Start(":12345")
}

//自定义上下文
type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Foo() {
	fmt.Println("foo")
}

func (c *CustomContext) Bar() {
	fmt.Println("bar")
}
