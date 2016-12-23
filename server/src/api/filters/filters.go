package filters

import (
	"../log"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func RegisterFilters() {
	log.Info("Filters initialize...")

	//beego.InsertFilter("*", beego.BeforeRouter, filer)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "UPDATE"},
		AllowCredentials: true,
	}))

	// register rbac filter
	//rbac.AccessRegister()
}
