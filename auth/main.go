package main

import (
	"github.com/astaxie/beego"
	_ "./conf/inits"
	_ "./routers"
)

func main() {
	beego.Run()
}
