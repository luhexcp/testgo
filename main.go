package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
)

func main() {
	// 读取 conf 配置，初始化数据库
	maxIdle, _ := beego.AppConfig.Int("maxIdle")
	maxConn, _ := beego.AppConfig.Int("maxConn")
	datasource := beego.AppConfig.String("datasource")
	orm.RegisterDataBase("default", "mysql", datasource, maxIdle, maxConn)
	fmt.println("dkk")

	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}
