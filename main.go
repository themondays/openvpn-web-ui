package main

import (
	"github.com/themondays/openvpn-web-ui/lib"
	_ "github.com/themondays/openvpn-web-ui/routers"
	"github.com/astaxie/beego"
)

func main() {
	lib.AddFuncMaps()
	beego.Run()
}
