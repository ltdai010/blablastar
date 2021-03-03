package main

import (
	_ "blablastar/routers"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/v1/bla-bla-star/swagger"] = "swagger"
	}
	web.Run()
}
