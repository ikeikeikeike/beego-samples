package main

import (
	"github.com/astaxie/beego"
	_ "github.com/ikeikeikeike/beego-samples/auth/conf/inits"
	_ "github.com/ikeikeikeike/beego-samples/auth/routers"
)

func main() {
	beego.Run()
}
